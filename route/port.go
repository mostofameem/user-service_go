package route

import (
	"context"
)

type Service interface {
	Register(ctx context.Context, user User) error
	Login(ctx context.Context, email, pass string) error
	Validate(ctx context.Context, email, code string) error
	Update(ctx context.Context, user User) error
	GetProfile(ctx context.Context, id int) (User, error)
}
