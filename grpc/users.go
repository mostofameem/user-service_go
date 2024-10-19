package grpc

import (
	"context"
	"log/slog"
	pb "user-service/grpc/users"
	"user-service/logger"
	"user-service/route"
)

type UserService struct {
	pb.UnimplementedUsersServer
	routeSvc route.Service
}

func NewUserService(routeSvc route.Service) *UserService {
	return &UserService{
		routeSvc: routeSvc,
	}
}

func (r *UserService) GetUserName(
	ctx context.Context,
	req *pb.GetUserNameReq,
) (*pb.GetUserNameRes, error) {
	userInfo, err := r.routeSvc.GetProfile(ctx, int(req.Id))
	if err != nil {
		slog.Error(err.Error(), logger.Extra(map[string]interface{}{
			"request": req,
		}))
		return nil, err
	}
	return &pb.GetUserNameRes{
		Status: true,
		Id:     int32(userInfo.Id),
		Name:   userInfo.Name,
	}, nil
}
