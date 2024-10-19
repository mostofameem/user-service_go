package route

import (
	"user-service/config"
	"user-service/db"
)

type service struct {
	cnf      *config.Config
	userRepo *db.UserTypeRepo
}

func NewService(cnf *config.Config) Service {
	return &service{
		cnf:      cnf,
		userRepo: db.NewUserTypeRepo(&cnf.DB),
	}
}
