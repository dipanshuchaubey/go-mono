package interfaces

import (
	us "carthage/protos/user_service"
	"context"
)

type UserOpsInterface interface {
	GetUsers(ctx context.Context, tenantID string) ([]*us.UserInfo, error)
	GetUserByID(ctx context.Context, tenantID, userIDs string) (*us.UserInfo, error)
}
