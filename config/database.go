package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DefaultTimeout = 30 * time.Second
	DefaultConfigID = 1
)

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