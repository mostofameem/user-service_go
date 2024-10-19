package route

import (
	"context"
	"time"
	"user-service/db"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Dob      string
	Type     string
}

func (svc *service) Register(ctx context.Context, user User) error {
	err := svc.userRepo.Register(&ctx, db.User{
		Name:       user.Name,
		Email:      user.Email,
		Pass:       user.Password,
		Dob:        user.Dob,
		Type:       user.Type,
		Is_active:  "false",
		Created_at: time.Now(),
		Updated_at: time.Now(),
	})

	return err
}
