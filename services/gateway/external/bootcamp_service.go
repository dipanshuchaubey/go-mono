package external

import (
	v1 "carthage/protos/bootcamp_service"
	pbrq "carthage/protos/bootcamp_service/request"
	pbrs "carthage/protos/bootcamp_service/response"
	"carthage/services/gateway/constants"
	"carthage/services/gateway/types"
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BootcampServiceInterface interface {
	GetBootcampsDetails(ctx context.Context, req *pbrq.GetBootcampsDetailsRequest) (*pbrs.GetBootcampsDetailsResponse, error)
	CreateBootcamp(ctx context.Context, req *pbrq.CreateBootcampRequest) (*pbrs.CreateBootcampResponse, error)
}

type BootcampServiceClient struct {
	bs     v1.BootcampServiceClient
	logger *slog.Logger
	tracer trace.Tracer
}

func BootcampService(config *types.Config) BootcampServiceInterface {
	conn, conErr := grpc.NewClient(config.EndPoints.GrpcBootcampService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if conErr != nil {
		fmt.Printf("Error connecting on bootcamp service: %v\n", conErr.Error())
	}

	c := v1.NewBootcampServiceClient(conn)

	return &BootcampServiceClient{
		bs:     c,
		logger: otelslog.NewLogger(constants.ServiceName),
		tracer: otel.Tracer(constants.ServiceName),
	}
}

func (b *BootcampServiceClient) GetBootcampsDetails(ctx context.Context, req *pbrq.GetBootcampsDetailsRequest) (*pbrs.GetBootcampsDetailsResponse, error) {
	b.logger.InfoContext(ctx, fmt.Sprintf("Calling User Service GetBootcampsDetails: %v", req))

	res, err := b.bs.GetBootcampsDetails(ctx, req)
	if err != nil {
		errMsg := fmt.Errorf("error calling GetBootcampDetails for BootcampIDs %s: %v", req.GetBootcampIds(), err)
		b.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	return res, nil
}

func (b *BootcampServiceClient) CreateBootcamp(ctx context.Context, req *pbrq.CreateBootcampRequest) (*pbrs.CreateBootcampResponse, error) {
	b.logger.InfoContext(ctx, fmt.Sprintf("Calling User Service CreateBootcamp: %v", req))

	res, err := b.bs.CreateBootcamp(ctx, req)
	if err != nil {
		errMsg := fmt.Errorf("error calling CreateBootcamp: %v", err)
		b.logger.ErrorContext(ctx, errMsg.Error())
		return nil, errMsg
	}

	return res, nil
}
