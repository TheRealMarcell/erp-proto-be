package usecases

import (
	"context"
	"erp-api/internal/modules/inventory"
	itemEntity "erp-api/internal/modules/item/models/entity"
	"erp-api/internal/modules/sale"
	"erp-api/internal/modules/transaction"
	"erp-api/internal/modules/transaction/models/request"
	"fmt"

	"erp-api/internal/pkg/helpers"
	"erp-api/internal/pkg/log"
)

type commandUsecase struct {
	transactionRepositoryCommand transaction.PostgresRepositoryCommand
	saleRepositoryCommand        sale.PostgresRepositoryCommand
	inventoryRepositoryCommand   inventory.PostgresRepositoryCommand
	saleRepositoryQuery          sale.PostgresRepositoryQuery
	logger                       log.Logger
}

func NewCommandUsecase(prq transaction.PostgresRepositoryCommand,
	sprc sale.PostgresRepositoryCommand,
	iprc inventory.PostgresRepositoryCommand,
	sprq sale.PostgresRepositoryQuery,
	log log.Logger) transaction.UsecaseCommand {
	return commandUsecase{
		transactionRepositoryCommand: prq,
		saleRepositoryCommand:        sprc,
		inventoryRepositoryCommand:   iprc,
		saleRepositoryQuery:          sprq,
		logger:                       log,
	}
}

func (c commandUsecase) InsertTransaction(ctx context.Context, payload request.Transaction) error {
	// save transaction and get the corresponding transaction id
	transactionId, err := c.transactionRepositoryCommand.SaveTransaction(ctx, payload)
	if err != nil {
		return err
	}

	// save sale
	if err := c.saleRepositoryCommand.BatchInsertSales(ctx, payload.Sales, transactionId); err != nil {
		return err
	}

	// sanitise location
	if err := helpers.IsValidLocation(payload.Location); err != nil {
		return err
	}

	// update stock in inventory (based on location)
	formatted_location := fmt.Sprintf("inventory_%v", payload.Location)

	var storageItems []itemEntity.StorageItem

	for _, sale := range payload.Sales {
		storageItem := itemEntity.StorageItem{
			Location:    sale.Location,
			ItemID:      sale.ItemID,
			Quantity:    sale.Quantity,
			Description: sale.Description,
		}
		storageItems = append(storageItems, storageItem)
	}

	if err := c.inventoryRepositoryCommand.BatchUpdateInventory(ctx, storageItems, formatted_location, "minus"); err != nil {
		return err
	}

	return nil
}

func (c commandUsecase) UpdatePaymentStatus(ctx context.Context, transactionId string, paymentStatus string) error {
	if err := c.transactionRepositoryCommand.ModifyPaymentStatus(ctx, transactionId, paymentStatus); err != nil {
		return err
	}
	return nil
}

func (c commandUsecase) DeleteTransaction(ctx context.Context, transactionId string) error {
	// TODO:⁠ ⁠get all sales based on transaction id provided
	// TODO: ⁠⁠update all quantities of each item_id in sales in the inventory based on location

	// 3.⁠ ⁠⁠delete the transaction (DONE)
	if err := c.transactionRepositoryCommand.RemoveTransaction(ctx, transactionId); err != nil {
		return err
	}

	return nil
}
