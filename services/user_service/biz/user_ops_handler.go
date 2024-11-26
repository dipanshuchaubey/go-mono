package biz

import (
	us "carthage/protos/user_service"
	"carthage/services/user_service/biz/interfaces"
	"carthage/services/user_service/types"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type UserOpsHandler struct{}

func NewUserOpsHandler() interfaces.UserOpsInterface {
	return &UserOpsHandler{}
}

func (h *UserOpsHandler) GetUsers(ctx context.Context, tenantID string) ([]*us.UserInfo, error) {
	// ctx, span := tracer.Start(ctx, "GetUsers")
	// defer span.End()

	httpRes, httpErr := http.Get("https://jsonplaceholder.typicode.com/users")

	if httpErr != nil {
		log.Fatalf("Cannot make http request %v", httpErr)
	}

	resBody, resErr := io.ReadAll(httpRes.Body)
	if resErr != nil {
		log.Fatal("Cannot read response body")
	}

	var userInfos []*types.UserInfo
	marErr := json.Unmarshal(resBody, &userInfos)

	if marErr != nil {
		log.Fatalf("Failed to unmarshal json response %v", marErr)
	}

	var response []*us.UserInfo
	for _, user := range userInfos {
		response = append(response, &us.UserInfo{
			UserId:   user.ID,
			Email:    user.Email,
			FullName: user.Name,
			UserType: us.UserTypes_USER_TYPE_ACTIVE,
		})
	}

	return response, nil
}

func (h *UserOpsHandler) GetUserByID(ctx context.Context, tenantID, userIDs string) (*us.UserInfo, error) {
	// ctx, span := tracer.Start(ctx, "GetUserByID")
	// defer span.End()

	httpRes, httpErr := http.Get("https://jsonplaceholder.typicode.com/users/" + userIDs)

	if httpErr != nil {
		log.Fatalf("Cannot make http request %v", httpErr)
	}

	resBody, resErr := io.ReadAll(httpRes.Body)
	if resErr != nil {
		log.Fatal("Cannot read response body")
	}

	var user *types.UserInfo
	marErr := json.Unmarshal(resBody, &user)

	if marErr != nil {
		log.Fatalf("Failed to unmarshal json response %v", marErr)
	}

	response := &us.UserInfo{
		UserId:   user.ID,
		Email:    user.Email,
		FullName: user.Name,
		UserType: us.UserTypes_USER_TYPE_ACTIVE,
	}

	return response, nil
}
