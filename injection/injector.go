package injection

import (
	"fmt"
	"oneclick/config"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	CacheInterval = 10 * time.Minute
)

type Injector struct {
	gormConnection *gorm.DB
	redisClient *redis.Client
}

func (i *Injector) ProvideGormConnection() *gorm.DB {
	if i.gormConnection == nil {
		l := zap.S()
		connection, err := config.NewCockroachDBConnection()
		if err != nil {
			l.Panic(err)
		}

		i.gormConnection = connection
	}

	return i.gormConnection
}

func (i *Injector) ProvideRedisClient() *redis.Client {
	var env config.Env

	env.LoadConfig()
	if i.redisClient == nil {
		addr := fmt.Sprintf("%v:%v", env.RedisHost, env.RedisPort)

		redisClient := redis.NewClient(&redis.Options{
			Addr: addr,
			Password: env.RedisPassword,
		})
		return redisClient
	}
	return i.redisClient
}