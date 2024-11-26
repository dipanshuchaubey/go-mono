package handlers

import (
	pb "carthage/protos/user_service"
	"carthage/services/gateway/constants"
	"carthage/services/gateway/external"
	"carthage/services/gateway/routes"
	"context"
	"fmt"
	"net/http"
)

type HandlerInterface interface {
	GetUsers() routes.HandlerFunc
	GetUser() routes.HandlerFunc
}

type userHandler struct {
	us external.UserServiceInterface
}

func UserHandler() HandlerInterface {
	us := external.UserService()
	return &userHandler{us}
}

func (h *userHandler) GetUsers() routes.HandlerFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		users, userErr := h.us.GetUsers(ctx, &pb.GetUsersRequest{TenantId: ""})
		if userErr != nil {
			return nil, userErr
		}

		return users.Data, nil
	}
}

func (h *userHandler) GetUser() routes.HandlerFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		userID := ctx.Value("id").(string)

		if userID == constants.EmptyString {
			errMsg := fmt.Errorf("user id not found")
			fmt.Println(errMsg)
			return nil, errMsg
		}

		user, userErr := h.us.GetUser(ctx, &pb.GetUserRequest{UserId: userID})
		if userErr != nil {
			return nil, userErr
		}

		return user.Data, nil
	}
}