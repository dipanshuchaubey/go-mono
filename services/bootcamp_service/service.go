package service

import (
	bootcamp "carthage/services/bootcamp_service/biz"
	"carthage/services/bootcamp_service/biz/interfaces"
	"carthage/services/bootcamp_service/config"
	"carthage/services/bootcamp_service/constants"
	"carthage/services/bootcamp_service/dto"
	"context"
	"fmt"
	"log/slog"

	pbrq "github.com/dipanshuchaubey/protos-package/bootcamp_service/request"
	pbrs "github.com/dipanshuchaubey/protos-package/bootcamp_service/response"

	v1 "github.com/dipanshuchaubey/protos-package/bootcamp_service"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type BootcampService struct {
	v1.UnimplementedBootcampServiceServer
	handler interfaces.BootcampInterface
	logger  *slog.Logger
	trace   trace.Tracer
}

func NewBootcampService(config config.Config) *BootcampService {
	return &BootcampService{
		UnimplementedBootcampServiceServer: v1.UnimplementedBootcampServiceServer{},
		handler:                            bootcamp.NewBootcampHandler(config),
		logger:                             otelslog.NewLogger(constants.ServiceName),
		trace:                              otel.Tracer(constants.ServiceName),
	}
}

func (s *BootcampService) GetBootcampsDetails(ctx context.Context, in *pbrq.GetBootcampsDetailsRequest) (*pbrs.GetBootcampsDetailsResponse, error) {
	s.logger.InfoContext(ctx, "Received GetBootcampsDetails request")

	res, err := s.handler.GetBootcampsDetails(ctx)
	if err != nil {
		return nil, err
	}

	return &pbrs.GetBootcampsDetailsResponse{Data: res}, nil
}

func (s *BootcampService) CreateBootcamp(ctx context.Context, in *pbrq.CreateBootcampRequest) (*pbrs.CreateBootcampResponse, error) {
	s.logger.InfoContext(ctx, fmt.Sprintf("Received CreateBootcamp request: %v", in))

	var body dto.CreateBootcampBody
	body.FromProto(in)

	res, err := s.handler.CreateBootcamp(ctx, body)
	if err != nil {
		return nil, err
	}

	return &pbrs.CreateBootcampResponse{Data: res, Success: true}, nil
}
