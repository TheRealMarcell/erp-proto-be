package usecases

import (
	"context"
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/inventory/models/request"
	"erp-api/internal/modules/item/models/entity"
	"erp-api/internal/pkg/log"
	"fmt"
)

type commandUsecase struct {
	inventoryRepositoryCommand inventory.PostgresRepositoryCommand
	logger                     log.Logger
}

func NewCommandUsecase(prc inventory.PostgresRepositoryCommand, log log.Logger) inventory.UsecaseCommand {
	return commandUsecase{
		inventoryRepositoryCommand: prc,
		logger:                     log,
	}
}

func (c commandUsecase) MoveInventory(ctx context.Context, payload request.MoveInventory) error {
	var storageItems []entity.StorageItem

	for _, i := range payload.Items {
		storageItem := entity.StorageItem{
			ItemID:   i.ItemID,
			Quantity: i.Quantity,
		}
		storageItems = append(storageItems, storageItem)
	}
	// for each item decrease qty in source and add to destination
	formatted_source := fmt.Sprintf("inventory_%v", payload.Source)
	formatted_destination := fmt.Sprintf("inventory_%v", payload.Destination)

	// batch update decrease
	if err := c.inventoryRepositoryCommand.BatchUpdateInventory(ctx, storageItems, formatted_source, "minus"); err != nil {
		return err
	}

	// batch update add
	if err := c.inventoryRepositoryCommand.BatchUpdateInventory(ctx, storageItems, formatted_destination, "add"); err != nil {
		return err
	}

	return nil
}
