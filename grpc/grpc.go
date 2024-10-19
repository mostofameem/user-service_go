package grpc

import (
	"sync"
	"user-service/config"
	"user-service/route"
)

type grpc struct {
	cnf      *config.Config
	routeSvc route.Service
	Wg       sync.WaitGroup
}

func NewGRPC(cnf *config.Config, routeSvc route.Service) *grpc {
	return &grpc{
		cnf:      cnf,
		routeSvc: routeSvc,
	}
}
