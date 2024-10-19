package redis

import (
	"context"
	"log/slog"
	"user-service/logger"

	"github.com/redis/go-redis/v9"
)

func GetCode(ctx context.Context, email string) (string, error) {
	code, err := Client.Get(ctx, email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		slog.Error(
			"Failed to retrieve secret code from Redis",
			logger.Extra(map[string]any{
				"error": err.Error(),
				"email": email,
			}),
		)
		return "", err
	}
	return code, nil
}
