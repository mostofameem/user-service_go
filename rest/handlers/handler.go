package handlers

import (
	"user-service/config"
	"user-service/route"
)

type Handlers struct {
	cnf      *config.Config
	routeSvc route.Service
}

func NewHandler(cnf *config.Config, routeSvc route.Service) *Handlers {
	return &Handlers{
		cnf:      cnf,
		routeSvc: routeSvc,
	}
}
