package external

import (
	v1 "carthage/protos/bootcamp_service"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BootcampServiceInterface interface {
	GetBootcampsDetails(ctx context.Context, req *v1.GetBootcampsDetailsRequest) (*v1.GetBootcampsDetailsResponse, error)
}

type BootcampServiceClient struct {
	bs v1.BootcampServiceClient
}

func BootcampService() BootcampServiceInterface {
	conn, conErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if conErr != nil {
		fmt.Printf("Error connecting on bootcamp service: %v\n", conErr.Error())
	}

	c := v1.NewBootcampServiceClient(conn)

	return &BootcampServiceClient{c}
}

func (b *BootcampServiceClient) GetBootcampsDetails(ctx context.Context, req *v1.GetBootcampsDetailsRequest) (*v1.GetBootcampsDetailsResponse, error) {
	fmt.Println("Calling User Service with: ", req)

	res, err := b.bs.GetBootcampsDetails(ctx, &v1.GetBootcampsDetailsRequest{})
	if err != nil {
		errMsg := fmt.Errorf("error calling GetBootcampDetails for BootcampIDs %s: %v", req.GetBootcampIds(), err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return res, nil
}
