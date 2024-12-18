package handlers

import (
	pbrq "carthage/protos/bootcamp_service/request"
	"carthage/services/gateway/external"
	"carthage/services/gateway/types"
	"context"
	"fmt"
	"io"
	"net/http"
)

type BootcampHandlerInterface interface {
	GetBootcamps() types.HandlerFunc
	CreateBootcamp() types.HandlerFunc
}

type bootcampHandler struct {
	bs external.BootcampServiceInterface
}

func BootcampHandler(config *types.Config) BootcampHandlerInterface {
	bs := external.BootcampService(config)
	return &bootcampHandler{bs}
}

func (h *bootcampHandler) GetBootcamps() types.HandlerFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		bootcamps, bootcampErr := h.bs.GetBootcampsDetails(ctx, &pbrq.GetBootcampsDetailsRequest{BootcampIds: []string{""}})
		if bootcampErr != nil {
			return nil, bootcampErr
		}

		return bootcamps.Data, nil
	}
}

func (h *bootcampHandler) CreateBootcamp() types.HandlerFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		body, _ := io.ReadAll(req.Body)
		fmt.Println("Creating Bootcamp: ", string(body))

		bootcamp, bootcampErr := h.bs.CreateBootcamp(ctx, &pbrq.CreateBootcampRequest{})
		if bootcampErr != nil {
			return nil, bootcampErr
		}

		return bootcamp.Data, nil
	}
}
