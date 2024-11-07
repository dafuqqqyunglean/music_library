package repository

import (
	"fmt"
	"github.com/dafuqqqyunglean/music_library/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func NewPostgresDB(cfg config.PostgresConfig, logger *zap.SugaredLogger) *sqlx.DB {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Fatalf("failed to connect to postgres: %v", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("failed to ping database: %v", err)
	}

	return db
}
