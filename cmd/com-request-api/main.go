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
