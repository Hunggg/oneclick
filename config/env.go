package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DatabaseUser      string
	DatabasePassword  string
	DatabaseHost      string
	DatabasePort      string
	DatabaseSchema    string
	DatabaseEnableLog string
	DatabaseEnableTLS string
	GrpcPort          string
	HttpPort          string
	EnableCors        string

	LogLevel  string
	EnableLog string

	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func (env *Env) LoadConfig() {
	viper.SetConfigFile("local.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	env.DatabaseUser = viper.GetString("DatabaseUser")
	env.DatabasePassword = viper.GetString("DatabasePassword")
	env.DatabaseHost = viper.GetString("DatabaseHost")
	env.DatabasePort = viper.GetString("DatabasePort")
	env.DatabaseSchema = viper.GetString("DatabaseSchema")
	env.DatabaseEnableLog = viper.GetString("DatabaseEnableLog")
	env.DatabaseEnableTLS = viper.GetString("DatabaseEnableTLS")
	env.GrpcPort = viper.GetString("GrpcPort")
	env.HttpPort = viper.GetString("HttpPort")
	env.EnableCors = viper.GetString("EnableCors")
	env.LogLevel = viper.GetString("LogLevel")
	env.EnableLog = viper.GetString("EnableLog")
	env.RedisHost = viper.GetString("RedisHost")
	env.RedisPort = viper.GetString("RedisPort")
	env.RedisPassword = viper.GetString("RedisPassword")
}
