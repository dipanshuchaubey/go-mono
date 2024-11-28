package service

import (
	bs "carthage/protos/bootcamp_service"
	bootcamp "carthage/services/bootcamp_service/biz"
	"carthage/services/bootcamp_service/biz/interfaces"
	"carthage/services/bootcamp_service/config"
	"carthage/services/bootcamp_service/dto"
	"context"
	"fmt"
)

type BootcampService struct {
	bs.UnimplementedBootcampServiceServer
	handler interfaces.BootcampInterface
}

func NewBootcampService(config config.Config) *BootcampService {
	return &BootcampService{
		UnimplementedBootcampServiceServer: bs.UnimplementedBootcampServiceServer{},
		handler:                            bootcamp.NewBootcampHandler(config),
	}
}

func (s *BootcampService) GetBootcampsDetails(ctx context.Context, in *bs.GetBootcampsDetailsRequest) (*bs.GetBootcampsDetailsResponse, error) {
	fmt.Println("Received GetBootcampsDetails request")

	res, err := s.handler.GetBootcampsDetails(ctx)
	if err != nil {
		errMsg := fmt.Errorf("error getting bootcamps %v", err)
		fmt.Println(errMsg)
		return nil, errMsg
	}

	return &bs.GetBootcampsDetailsResponse{Data: res}, nil
}

func (s *BootcampService) CreateBootcamp(ctx context.Context, in *bs.CreateBootcampRequest) (*bs.CreateBootcampResponse, error) {
	fmt.Printf("Received CreateBootcamp request: %v\n", in)

	var body dto.CreateBootcampBody
	body.FromProto(in)

	res, err := s.handler.CreateBootcamp(ctx, body)
	if err != nil {
		return nil, err
	}

	return &bs.CreateBootcampResponse{Data: res, Success: true}, nil
}
