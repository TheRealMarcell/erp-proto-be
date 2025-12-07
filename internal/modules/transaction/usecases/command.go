package usecases

import (
	"context"
	"erp-api/internal/modules/inventory"
	itemEntity "erp-api/internal/modules/item/models/entity"
	"erp-api/internal/modules/sale"
	"erp-api/internal/modules/sale/models/entity"
	"erp-api/internal/modules/transaction"
	"erp-api/internal/modules/transaction/models/request"
	"fmt"

	"erp-api/internal/pkg/errors"
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
	respSales := <-c.saleRepositoryQuery.FindAllSales(ctx, transactionId)
	if respSales.Error != nil {
		msg := "No sales found"
		return errors.NotFound(msg)
	}

	salesData, ok := respSales.Data.([]entity.Sale)

	if !ok {
		return errors.InternalServerError("Cannot parse sales data")
	}

	var items []itemEntity.StorageItem

	for _, sale := range salesData {
		item := itemEntity.StorageItem{
			Location:    sale.Location,
			ItemID:      sale.ItemID,
			Quantity:    sale.Quantity,
			Description: "",
		}

		items = append(items, item)
	}

	if err := c.inventoryRepositoryCommand.BatchUpdateInventory(ctx, items, "", "add"); err != nil {
		return err
	}

	if err := c.saleRepositoryCommand.BatchDeleteSales(ctx, salesData); err != nil {
		return err
	}

	if err := c.transactionRepositoryCommand.RemoveTransaction(ctx, transactionId); err != nil {
		return err
	}

	return nil
}
