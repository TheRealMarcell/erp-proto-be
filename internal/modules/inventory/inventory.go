package inventory

import (
	"context"
	"erp-api/internal/modules/inventory/models/response"
	itemEntity "erp-api/internal/modules/item/models/entity"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetInventory(ctx context.Context, location string) ([]response.InventoryData, error)
}

type UsecaseCommand interface {
}

type PostgresRepositoryQuery interface {
	FindListInventory(ctx context.Context, location string) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	BatchInsertInventory(ctx context.Context, items []itemEntity.Item) error
	BatchUpdateInventory(ctx context.Context, storageItems []itemEntity.StorageItem, location string, operation string) error

	UpdateInventory(ctx context.Context, storageItem itemEntity.StorageItem) error
}
