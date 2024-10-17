package rabbitmq

import (
	"log/slog"
	"os"
	"user-service/rabbitmq/consumers"
)

func RunConsumers() {
	client := GetClient()
	if client == nil {
		slog.Error("No client")
		os.Exit(1)
		return
	}

	// -------------------------------------- User --------------------------------
	client.AddConsumer(
		ConsumerOption{Exchange: ExchangeName,
			RoutingKey: consumers.UserRoutingKey,
			Queue:      consumers.UserQueueName(),
			Consumer:   consumers.ConsumeUser,
		},
	)

}
