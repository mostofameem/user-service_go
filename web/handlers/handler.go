package handlers

import "base_service/config"

type Handlers struct {
	cnf *config.Config
}

func NewHandler(cnf *config.Config) *Handlers {
	return &Handlers{
		cnf: cnf,
	}
}
