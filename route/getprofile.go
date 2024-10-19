package route

import (
	"context"
	"log/slog"
)

func (svc *service) GetProfile(ctx context.Context, id int) (*User, error) {
	user, err := svc.userRepo.GetProfile(ctx, id)
	if err != nil {
		slog.Error("failed to get user info", err)
		return nil, err
	}
	
	return &User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Type:  user.Type,
	}, nil
}
