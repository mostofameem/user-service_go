package redis

import (
	"log/slog"
	"os"
	"sync"
	"user-service/config"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

var redisCntOnce = sync.Once{}

func NewRedis(cnf *config.Config) *redis.Client {
	redisCntOnce.Do(func() {
		url, err := redis.ParseURL(cnf.RedisURL)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
			return
		}
		Client = redis.NewClient(url)
	})
	return Client
}
