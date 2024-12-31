package interfaces

import (
	"carthage/services/bootcamp_service/dto"
	"context"

	pbrs "github.com/dipanshuchaubey/protos-package/bootcamp_service/response"
	pbty "github.com/dipanshuchaubey/protos-package/bootcamp_service/types"
)

type BootcampInterface interface {
	GetBootcampsDetails(ctx context.Context) ([]*pbrs.GetBootcampsDetailsResponse_Data, error)
	CreateBootcamp(ctx context.Context, body dto.CreateBootcampBody) (*pbty.BootcampInfo, error)
}
