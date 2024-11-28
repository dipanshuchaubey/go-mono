package interfaces

import (
	bs "carthage/protos/bootcamp_service"
	"carthage/services/bootcamp_service/dto"
	"context"
)

type BootcampInterface interface {
	GetBootcampsDetails(ctx context.Context) ([]*bs.GetBootcampsDetailsResponse_Data, error)
	CreateBootcamp(ctx context.Context, body dto.CreateBootcampBody) (*bs.BootcampInfo, error)
}
