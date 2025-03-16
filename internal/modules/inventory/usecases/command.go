package usecases

import (
	"context"
	"erp-api/internal/modules/history"
	historyEntity "erp-api/internal/modules/history/models/entity"
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/inventory/models/request"
	itemEntity "erp-api/internal/modules/item/models/entity"
	"erp-api/internal/pkg/helpers"
	"erp-api/internal/pkg/log"
	"errors"
	"fmt"
	"time"
)

type commandUsecase struct {
	inventoryRepositoryCommand inventory.PostgresRepositoryCommand
	historyRepositoryCommand   history.PostgresRepositoryCommand
	logger                     log.Logger
}

func NewCommandUsecase(prc inventory.PostgresRepositoryCommand, hprc history.PostgresRepositoryCommand, log log.Logger) inventory.UsecaseCommand {
	return commandUsecase{
		inventoryRepositoryCommand: prc,
		historyRepositoryCommand:   hprc,
		logger:                     log,
	}
}

func (c commandUsecase) MoveInventory(ctx context.Context, payload request.MoveInventory) error {
	var storageItems []itemEntity.StorageItem

	for _, i := range payload.Items {
		storageItem := itemEntity.StorageItem{
			ItemID:   i.ItemID,
			Quantity: i.Quantity,
		}
		storageItems = append(storageItems, storageItem)
	}

	// sanitise inventory locations
	if err := helpers.IsValidLocation(payload.Source); err != nil {
		return err
	}

	if err := helpers.IsValidLocation(payload.Destination); err != nil {
		return err
	}

	if payload.Destination == payload.Source {
		return errors.New("source and destionation cannot be the same")
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

	// batch insert into history
	var historyItems []historyEntity.History

	for _, i := range payload.Items {
		currentTime := time.Now()

		historyItem := historyEntity.History{
			ItemID:      i.ItemID,
			Quantity:    i.Quantity,
			Timestamp:   currentTime,
			Source:      payload.Source,
			Destination: payload.Destination,
		}
		historyItems = append(historyItems, historyItem)
	}

	if err := c.historyRepositoryCommand.BatchInsertHistory(ctx, historyItems); err != nil {
		return err
	}

	return nil
}
