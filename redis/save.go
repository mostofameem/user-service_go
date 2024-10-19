package redis

import (
	"context"
	"log/slog"
	"time"
	"user-service/logger"
)

func Save(ctx context.Context, email string, secretCode string, duration int) error {
	err := Client.Set(ctx, email, secretCode, time.Duration(duration)*time.Minute).Err()
	if err != nil {
		slog.Error(
			"Failed to store secret code in Redis",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"email": email,
			}),
		)
		return err
	}
	return nil
}
