package database

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
)

// NewPostgres returns DB
func NewPostgres(dsn, driver string, retry int) (*sqlx.DB, error) {
	var err error
	var db *sqlx.DB
	for retry > 0 {
		db, err = sqlx.Open(driver, dsn)
		if err != nil {
			log.Error().Err(err).Msgf("failed to create database connection")
			retry--
			time.Sleep(3 * time.Second)
			continue
		}

		if err = db.Ping(); err != nil {
			log.Error().Err(err).Msgf("failed ping the database")
			retry--
			time.Sleep(3 * time.Second)
			continue
		}

		return db, nil
	}
	return db, err
}
