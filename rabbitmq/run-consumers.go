package rabbitmq

import (
	"log/slog"
	"base_service/rabbitmq/consumers"
	"os"
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
