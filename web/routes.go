package web

import (
	"base_service/web/middlewares"
	"net/http"
)

func (server *Server) initRouts(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

}
