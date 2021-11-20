package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/denlipov/com-request-api/internal/app/repo"
	"github.com/denlipov/com-request-api/internal/app/retranslator"
	"github.com/denlipov/com-request-api/internal/app/sender"
	"github.com/denlipov/com-request-api/internal/config"
	"github.com/denlipov/com-request-api/internal/database"
	"github.com/denlipov/com-request-api/internal/tracer"
	"github.com/halink0803/zerolog-graylog-hook/graylog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	hook, err := graylog.NewGraylogHook(
		fmt.Sprintf("%s://%s:%d",
			cfg.Graylog.Proto,
			cfg.Graylog.Host,
			cfg.Graylog.Port))
	if err != nil {
		log.Error().Msgf("Unable to connect to graylog service: %+v", err)
	} else {
		//Set global logger with graylog hook
		log.Logger = log.Hook(hook)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	log.Debug().Msg("Db initialized")

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed init tracing")
		return
	}
	defer tracing.Close()

	sigs := make(chan os.Signal, 1)

	xlatorCfg := retranslator.Config{
		ChannelSize:    128,
		ConsumerCount:  10,
		ConsumeTimeout: 1 * time.Second,
		ConsumeSize:    10,
		ProducerCount:  10,
		WorkerCount:    4,
		Repo:           repo.NewEventRepo(db),
		Sender:         sender.NewEventSender(),
	}

	retranslator := retranslator.NewRetranslator(xlatorCfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	retranslator.Close()
}
