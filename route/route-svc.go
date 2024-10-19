package route

import (
	"sync"
	"user-service/config"
	"user-service/db"
	"user-service/redis"

	goRedis "github.com/redis/go-redis/v9"
)

type service struct {
	cnf      *config.Config
	userRepo *db.UserTypeRepo
	redis    *goRedis.Client
}

var routeSvc *service

var routeSvcCnt = sync.Once{}

func NewService(cnf *config.Config) Service {
	routeSvcCnt.Do(func() {
		routeSvc = &service{
			cnf:      cnf,
			userRepo: db.NewUserTypeRepo(&cnf.DB),
			redis:    redis.NewRedis(cnf),
		}
	})

	return routeSvc
}
