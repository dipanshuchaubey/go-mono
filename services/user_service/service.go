package service

import (
	us "carthage/protos/user_service"
	"carthage/services/user_service/biz"
	"carthage/services/user_service/biz/interfaces"
	"carthage/services/user_service/constants"
	"context"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

type UserService struct {
	us.UnimplementedUserServiceServer
	handler interfaces.UserOpsInterface
	logger  *slog.Logger
	tracer  trace.Tracer
}

func NewUserService() *UserService {
	return &UserService{
		UnimplementedUserServiceServer: us.UnimplementedUserServiceServer{},
		handler:                        biz.NewUserOpsHandler(),
		logger:                         otelslog.NewLogger(constants.ServiceName),
		tracer:                         otel.Tracer(constants.ServiceName),
	}
}

func (h *UserService) GetUsers(ctx context.Context, in *us.GetUsersRequest) (*us.GetUsersResponse, error) {
	ctx, span := h.tracer.Start(ctx, "service.GetUsers")
	defer span.End()

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
	ctx, span := h.tracer.Start(ctx, "service.GetUserByID")
	defer span.End()

	response, err := h.handler.GetUserByID(ctx, in.TenantId, in.UserId)
	if err != nil {
		return nil, err
	}

	return &us.GetUserResponse{
		Success: true,
		Data:    response,
	}, nil
}
