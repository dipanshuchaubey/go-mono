package handlers

import (
	pb "carthage/protos/bootcamp_service"
	"carthage/services/gateway/external"
	"carthage/services/gateway/routes"
	"context"
	"net/http"
)

type BootcampHandlerInterface interface {
	GetBootcamps() routes.HandlerFunc
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
