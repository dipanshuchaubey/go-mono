package external

import (
	v1 "carthage/protos/bootcamp_service"
	pbrq "carthage/protos/bootcamp_service/request"
	pbrs "carthage/protos/bootcamp_service/response"
	"carthage/services/gateway/types"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BootcampServiceInterface interface {
	GetBootcampsDetails(ctx context.Context, req *pbrq.GetBootcampsDetailsRequest) (*pbrs.GetBootcampsDetailsResponse, error)
	CreateBootcamp(ctx context.Context, req *pbrq.CreateBootcampRequest) (*pbrs.CreateBootcampResponse, error)
}

type BootcampServiceClient struct {
	bs v1.BootcampServiceClient
}

func BootcampService(config *types.Config) BootcampServiceInterface {
	conn, conErr := grpc.NewClient(config.EndPoints.GrpcBootcampService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if conErr != nil {
		fmt.Printf("Error connecting on bootcamp service: %v\n", conErr.Error())
	}

	c := v1.NewBootcampServiceClient(conn)

	return &BootcampServiceClient{c}
}

func (b *BootcampServiceClient) GetBootcampsDetails(ctx context.Context, req *pbrq.GetBootcampsDetailsRequest) (*pbrs.GetBootcampsDetailsResponse, error) {
	fmt.Println("Calling User Service GetBootcampsDetails: ", req)

	res, err := b.bs.GetBootcampsDetails(ctx, req)
	if err != nil {
		errMsg := fmt.Errorf("error calling GetBootcampDetails for BootcampIDs %s: %v", req.GetBootcampIds(), err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return res, nil
}

func (b *BootcampServiceClient) CreateBootcamp(ctx context.Context, req *pbrq.CreateBootcampRequest) (*pbrs.CreateBootcampResponse, error) {
	fmt.Println("Calling User Service CreateBootcamp: ", req)

	res, err := b.bs.CreateBootcamp(ctx, req)
	if err != nil {
		errMsg := fmt.Errorf("error calling CreateBootcamp: %v", err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return res, nil
}
