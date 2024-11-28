package handlers

import (
	pb "carthage/protos/bootcamp_service"
	"carthage/services/gateway/external"
	"carthage/services/gateway/routes"
	"context"
	"fmt"
	"io"
	"net/http"
)

type BootcampHandlerInterface interface {
	GetBootcamps() routes.HandlerFunc
	CreateBootcamp() routes.HandlerFunc
}

type bootcampHandler struct {
	bs external.BootcampServiceInterface
}

func BootcampHandler() BootcampHandlerInterface {
	bs := external.BootcampService()
	return &bootcampHandler{bs}
}

func (h *bootcampHandler) GetBootcamps() routes.HandlerFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		bootcamps, bootcampErr := h.bs.GetBootcampsDetails(ctx, &pb.GetBootcampsDetailsRequest{BootcampIds: []string{""}})
		if bootcampErr != nil {
			return nil, bootcampErr
		}

		return bootcamps.Data, nil
	}
}

func (h *bootcampHandler) CreateBootcamp() routes.HandlerFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		body, _ := io.ReadAll(req.Body)
		fmt.Println("Creating Bootcamp: ", string(body))

		bootcamp, bootcampErr := h.bs.CreateBootcamp(ctx, &pb.CreateBootcampRequest{})
		if bootcampErr != nil {
			return nil, bootcampErr
		}

		return bootcamp.Data, nil
	}
}
