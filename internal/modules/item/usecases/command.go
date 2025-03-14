package usecases

import (
	"context"
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/entity"
	"erp-api/internal/modules/item/models/request"
	"erp-api/internal/modules/sale"
	"erp-api/internal/pkg/log"
)

type commandUsecase struct {
	itemRepositoryCommand      item.PostgresRepositoryCommand
	inventoryRepositoryCommand inventory.PostgresRepositoryCommand
	saleRepositoryCommand      sale.PostgresRepositoryCommand
	logger                     log.Logger
}

func NewCommandUsecase(
	prq item.PostgresRepositoryCommand,
	iprq inventory.PostgresRepositoryCommand,
	sprq sale.PostgresRepositoryCommand,
	log log.Logger) item.UsecaseCommand {
	return commandUsecase{
		itemRepositoryCommand:      prq,
		inventoryRepositoryCommand: iprq,
		saleRepositoryCommand:      sprq,
		logger:                     log,
	}
}

func (c commandUsecase) SaveItem(ctx context.Context, payload request.SubmitItem) error {
	// insert item into items repository
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
	var storageItems []entity.StorageItem

	for _, i := range payload.Items {
		storageItem := entity.StorageItem{
			ItemID:   i.ItemID,
			Quantity: i.Quantity,
		}
		storageItems = append(storageItems, storageItem)
	}

	// add qty in inventory gudang
	if err := c.inventoryRepositoryCommand.BatchUpdateAddInventory(ctx, storageItems, "inventory_gudang"); err != nil {
		return err
	}

	// update retur qty in sale
	if err := c.saleRepositoryCommand.BatchUpdateReturQty(ctx, payload); err != nil {
		return err
	}

	return nil
}

func (c commandUsecase) CorrectItem(ctx context.Context, payload request.CorrectItem, id string) error {
	storageItem := entity.StorageItem{
		ItemID:   id,
		Location: payload.Location,
		Quantity: payload.Quantity,
	}

	if err := c.inventoryRepositoryCommand.UpdateInventory(ctx, storageItem); err != nil {
		return err
	}

	return nil
}

func (c commandUsecase) BrokenItem(ctx context.Context, payload request.UpdateItem) error {
	var storageItems []entity.StorageItem

	for _, i := range payload.Items {
		storageItem := entity.StorageItem{
			ItemID:   i.ItemID,
			Quantity: i.Quantity,
		}
		storageItems = append(storageItems, storageItem)
	}

	// add qty in inventory rusak
	if err := c.inventoryRepositoryCommand.BatchUpdateAddInventory(ctx, storageItems, "inventory_rusak"); err != nil {
		return err
	}

	// update qty retur in inventory rusak
	if err := c.saleRepositoryCommand.BatchUpdateReturQty(ctx, payload); err != nil {
		return err
	}

	return nil
}

func (c commandUsecase) UpdateItemPrice(ctx context.Context, payload request.ItemPrice) error {
	if err := c.itemRepositoryCommand.ModifyItemPrice(ctx, payload); err != nil {
		return err
	}

	return nil
}
