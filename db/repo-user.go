package db

import (
	"sync"
	"user-service/config"

	"github.com/jmoiron/sqlx"
)

type UserTypeRepo struct {
	db    *sqlx.DB
	table string
}

var userTypeRepo *UserTypeRepo

var userCntOnce = sync.Once{}

func NewUserTypeRepo(cnf *config.DBConfig) *UserTypeRepo {
	userCntOnce.Do(func() {
		db := connect(cnf)
		userTypeRepo = &UserTypeRepo{
			db:    db,
			table: "users",
		}
	})
	return userTypeRepo
}
