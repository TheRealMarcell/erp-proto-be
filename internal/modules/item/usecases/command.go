package usecases

import (
	"context"
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/request"
	"erp-api/internal/pkg/log"
)

type commandUsecase struct {
	itemRepositoryCommand      item.PostgresRepositoryCommand
	inventoryRepositoryCommand inventory.PostgresRepositoryCommand
	logger                     log.Logger
}

func NewCommandUsecase(
	prq item.PostgresRepositoryCommand,
	iprq inventory.PostgresRepositoryCommand,
	log log.Logger) item.UsecaseCommand {
	return commandUsecase{
		itemRepositoryCommand:      prq,
		inventoryRepositoryCommand: iprq,
		logger:                     log,
	}
}

func (c commandUsecase) SaveItem(ctx context.Context, payload request.SubmitItem) error {
	// insert item into items table
	if err := c.itemRepositoryCommand.BatchInsertItems(ctx, payload.Items); err != nil {
		return err
	}

	// insert item into each inventory
	if err := c.inventoryRepositoryCommand.BatchInsertInventory(ctx, payload.Items); err != nil {
		return err
	}

	return nil
}

func (c commandUsecase) UpdateItem(ctx context.Context, payload request.UpdateItem) error {
	return nil
}
