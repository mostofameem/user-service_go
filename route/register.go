package route

import (
	"context"
	"time"
)

type User struct {
	Name     string
	Email    string
	Password string
	Dob      time.Time
	Type     string
}

func (svc *service) Register(ctx context.Context, user User) error {

	return nil
}
