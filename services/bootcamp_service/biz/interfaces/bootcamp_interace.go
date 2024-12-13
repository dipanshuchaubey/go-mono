package interfaces

import (
	pbrs "carthage/protos/bootcamp_service/response"
	pbty "carthage/protos/bootcamp_service/types"
	"carthage/services/bootcamp_service/dto"
	"context"
)

type BootcampInterface interface {
	GetBootcampsDetails(ctx context.Context) ([]*pbrs.GetBootcampsDetailsResponse_Data, error)
	CreateBootcamp(ctx context.Context, body dto.CreateBootcampBody) (*pbty.BootcampInfo, error)
}
