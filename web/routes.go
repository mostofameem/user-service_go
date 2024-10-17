package web

import (
	"net/http"
	"user-service/web/middlewares"
)

func (server *Server) initRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello-world",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

}
