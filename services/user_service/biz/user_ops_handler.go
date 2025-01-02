package biz

import (
	us "carthage/protos/user_service"
	"carthage/services/user_service/biz/interfaces"
	"carthage/services/user_service/constants"
	"carthage/services/user_service/types"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type UserOpsHandler struct {
	logger *slog.Logger
	tracer trace.Tracer
}

func NewUserOpsHandler() interfaces.UserOpsInterface {
	return &UserOpsHandler{
		logger: otelslog.NewLogger(constants.ServiceName),
		tracer: otel.Tracer(constants.ServiceName),
	}
}

func (h *UserOpsHandler) GetUsers(ctx context.Context, tenantID string) ([]*us.UserInfo, error) {
	ctx, span := h.tracer.Start(ctx, "biz.GetUsers")
	defer span.End()

	httpRes, httpErr := http.Get("https://jsonplaceholder.typicode.com/users")

	if httpErr != nil {
		errMsg := fmt.Errorf("cannot make http request %v", httpErr)
		h.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	resBody, resErr := io.ReadAll(httpRes.Body)
	if resErr != nil {
		errMsg := fmt.Errorf("cannot read response body %v", resErr)
		h.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	var userInfos []*types.UserInfo
	marErr := json.Unmarshal(resBody, &userInfos)

	if marErr != nil {
		errMsg := fmt.Errorf("failed to unmarshal json response %v", marErr)
		h.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
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
	ctx, span := h.tracer.Start(ctx, "biz.GetUserByID")
	defer span.End()

	httpRes, httpErr := http.Get("https://jsonplaceholder.typicode.com/users/" + userIDs)

	if httpErr != nil {
		errMsg := fmt.Errorf("cannot make http request %v", httpErr)
		h.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	resBody, resErr := io.ReadAll(httpRes.Body)
	if resErr != nil {
		errMsg := fmt.Errorf("cannot read response body %v", resErr)
		h.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	var user *types.UserInfo
	marErr := json.Unmarshal(resBody, &user)

	if marErr != nil {
		errMsg := fmt.Errorf("failed to unmarshal json response %v", marErr)
		h.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	response := &us.UserInfo{
		UserId:   user.ID,
		Email:    user.Email,
		FullName: user.Name,
		UserType: us.UserTypes_USER_TYPE_ACTIVE,
	}

	return response, nil
}
