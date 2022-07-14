package config

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DefaultTimeout = 30 * time.Second
	DefaultConfigID = 1
)

type CockroachDB struct {
	l *zap.SugaredLogger
	db *gorm.DB
}

func NewCockroachDB(db *gorm.DB) (*CockroachDB, error) {
	return &CockroachDB{
		l: zap.S(),
		db: db,
	}, nil
}

func NewCockroachDBConnection() (*gorm.DB, error) {
	var env Env
	env.LoadConfig()
	uriDB := "postgresql://"+ env.DatabaseUser + "@" + env.DatabaseHost + ":" + env.DatabasePort + env.DatabaseSchema + "?sslmode=disable"
	logLevel := logger.Silent

	switch env.LogLevel {
	case "info":
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(uriDB), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}