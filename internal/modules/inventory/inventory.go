package inventory

import (
	"context"
	"erp-api/internal/modules/inventory/models/response"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetInventory(ctx context.Context, location string) ([]response.InventoryData, error)
}

type UsecaseCommand interface {

}

type PostgresRepositoryQuery interface {
	GetInventory(ctx context.Context, location string) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	
}