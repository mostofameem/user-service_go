package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"user-service/config"
	"user-service/rest/handlers"
	"user-service/rest/middlewares"
	"user-service/rest/swagger"
)

type Server struct {
	handlers *handlers.Handlers
	cnf      *config.Config
	Wg       sync.WaitGroup
}

func NewServer(cnf *config.Config, handlers *handlers.Handlers) (*Server, error) {
	server := &Server{
		cnf:      cnf,
		handlers: handlers,
	}

	return server, nil
}

func (server *Server) Start() {
	manager := middlewares.NewManager()

	manager.Use(
		middlewares.Logger,
	)

	mux := http.NewServeMux()

	server.initRouts(mux, manager)

	handler := middlewares.EnableCors(mux)

	swagger.SetupSwagger(mux, manager)

	server.Wg.Add(1)

	go func() {
		defer server.Wg.Done()

		addr := fmt.Sprintf(":%d", server.cnf.HttpPort)
		slog.Info(fmt.Sprintf("Listening at %s", addr))

		if err := http.ListenAndServe(addr, handler); err != nil {
			slog.Error(err.Error())
		}
	}()
}
