package rest

import (
	"net/http"
	"user-service/rest/middlewares"
)

func (server *Server) initRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello-world",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

	mux.Handle(
		"POST /register",
		manager.With(
			http.HandlerFunc(server.handlers.Register),
		),
	)

}
