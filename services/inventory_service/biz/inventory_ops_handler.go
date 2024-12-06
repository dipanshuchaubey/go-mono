package biz

import (
	pb "carthage/protos/inventory_service"
)

type InventoryOpsHandler struct{}

type InventoryOpsInterface interface {
	ListInventoryWithFilters(tenantID string, filters InventoryFilters) ([]*pb.InventoryInfo, error)
}

func NewInventoryOpsHandler() InventoryOpsInterface {
	return &InventoryOpsHandler{}
}

type InventoryFilters struct {
	Status     string
	EntityType string
	PageSize   int32
}

func (h *InventoryOpsHandler) ListInventoryWithFilters(tenantID string, filters InventoryFilters) ([]*pb.InventoryInfo, error) {
	return []*pb.InventoryInfo{}, nil
}
