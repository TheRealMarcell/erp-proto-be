package usecases

import (
	"context"
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/inventory/models/entity"
	"erp-api/internal/modules/inventory/models/response"
	"erp-api/internal/pkg/errors"
	"erp-api/internal/pkg/log"
	"fmt"
)

type queryUsecase struct {
	inventoryRepositoryQuery inventory.PostgresRepositoryQuery
	logger log.Logger
}

func NewQueryUsecase(
	prq inventory.PostgresRepositoryQuery,
	log log.Logger) inventory.UsecaseQuery {
		return queryUsecase{
			inventoryRepositoryQuery: prq,
			logger: log,
		}
	}

func (q queryUsecase) GetInventory(ctx context.Context, location string) ([]response.InventoryData, error){
	respInventory := <- q.inventoryRepositoryQuery.GetInventory(ctx, location)
	if respInventory.Error != nil{
		msg := fmt.Sprintf("No inventory found for location: %v", location)
		return nil, errors.NotFound(msg)
	}

	inventoryData, ok := respInventory.Data.([]entity.InventoryItem)

	if !ok {
		return nil, errors.InternalServerError("Cannot parse inventory data")
	}

	var resp []response.InventoryData

	for _, inventoryItem := range inventoryData{
		item := response.InventoryData{
			ItemID:				inventoryItem.ItemID,
			Quantity:			inventoryItem.Quantity,
			Description:	inventoryItem.Description,
			Price:				inventoryItem.Price,
		}
		resp = append(resp, item)
	}

	return resp, nil
}