package service

import (
	us "carthage/protos/user_service"
	"carthage/services/user_service/biz"
	"carthage/services/user_service/biz/interfaces"
	"context"
	"fmt"
	"time"
)

type UserService struct {
	us.UnimplementedUserServiceServer
	handler interfaces.UserOpsInterface
}

func NewUserService() *UserService {
	return &UserService{
		UnimplementedUserServiceServer: us.UnimplementedUserServiceServer{},
		handler:                        biz.NewUserOpsHandler(),
	}
}

func (h *UserService) GetUsers(ctx context.Context, in *us.GetUsersRequest) (*us.GetUsersResponse, error) {
	// ctx, span := tracer.Start(ctx, "GetUsers")
	// defer span.End()

	response, err := h.handler.GetUsers(ctx, in.TenantId)
	if err != nil {
		return nil, err
	}

	return &us.GetUsersResponse{
		Success: true,
		Data:    response,
	}, nil
}

func (h *UserService) GetUser(ctx context.Context, in *us.GetUserRequest) (*us.GetUserResponse, error) {
	// ctx, span := tracer.Start(ctx, "GetUserByID")
	// defer span.End()

	deadline, _ := ctx.Deadline()
	fmt.Println("Time remaining: ", time.Until(deadline))

	response, err := h.handler.GetUserByID(ctx, in.TenantId, in.UserId)
	if err != nil {
		return nil, err
	}

	return &us.GetUserResponse{
		Success: true,
		Data:    response,
	}, nil
}
