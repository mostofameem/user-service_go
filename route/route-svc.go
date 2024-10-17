package route

import "user-service/config"

type service struct {
	cnf *config.Config
}

func NewService(cnf *config.Config) Service {
	return &service{
		cnf: cnf,
	}
}
