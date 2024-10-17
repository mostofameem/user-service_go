package main

import (
	"log/slog"
	"user-service/config"
	"user-service/logger"
	"user-service/web"
	"user-service/web/handlers"
)

func main() {
	cnf := config.GetConfig()

	logger.SetupLogger(cnf.ServiceName)

	handlers := handlers.NewHandler(cnf)

	server, err := web.NewServer(cnf, handlers)
	if err != nil {
		slog.Error("failed to create the server:", logger.Extra(map[string]any{
			"error": err,
		}))
		return
	}
	server.Start()
	server.Wg.Wait()
}
