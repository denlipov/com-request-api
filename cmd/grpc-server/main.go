package main

import (
	"flag"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/denlipov/com-request-api/internal/cache"
	"github.com/denlipov/com-request-api/internal/config"
	"github.com/denlipov/com-request-api/internal/database"
	"github.com/denlipov/com-request-api/internal/server"
	"github.com/denlipov/com-request-api/internal/tracer"
	"github.com/halink0803/zerolog-graylog-hook/graylog"
)

var (
	batchSize uint = 2
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	cfg := config.GetConfigInstance()

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
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

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver, cfg.Database.Retry)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			log.Error().Err(err).Msg("Migration failed")

			return
		}
	}

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed init tracing")

		return
	}
	defer tracing.Close()

	cache := cache.NewRedisCache(cfg.Redis)
	if err := server.NewGrpcServer(db, cache, batchSize).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
