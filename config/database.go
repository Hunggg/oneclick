package config

import (
	"oneclick/entity"
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
	tables := make([]interface{}, 0)
	tables = append(tables, &entity.Categories{})

	if err := db.AutoMigrate(tables...); err != nil {
		return nil, err
	}

	return &CockroachDB{
		l: zap.S(),
		db: db,
	}, nil
}

func NewCockroachDBConnection() (*gorm.DB, error) {
	var env Env
	env.LoadConfig()
	var pass string
	if env.DatabasePassword != "" {
		pass = ":" + env.DatabasePassword
	} else {
		pass = env.DatabasePassword
	}
	uriDB := "postgresql://"+ env.DatabaseUser + pass + "@" + env.DatabaseHost + ":" + env.DatabasePort + "/" + env.DatabaseSchema + "?sslmode=disable"
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