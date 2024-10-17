package main

import (
	"log/slog"
	"user-service/config"
	"user-service/logger"
	"user-service/rest"
	"user-service/rest/handlers"
	"user-service/route"
)

func main() {
	cnf := config.GetConfig()

	logger.SetupLogger(cnf.ServiceName)

	routeSvc := route.NewService(cnf)

	handlers := handlers.NewHandler(cnf, routeSvc)

	server, err := rest.NewServer(cnf, handlers)
	if err != nil {
		slog.Error("failed to create the server:", logger.Extra(map[string]any{
			"error": err,
		}))
		return
	}
	server.Start()
	server.Wg.Wait()
}
