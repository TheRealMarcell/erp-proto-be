package usecases

import (
	"context"
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/entity"
	"erp-api/internal/modules/item/models/response"
	"erp-api/internal/pkg/errors"
	"erp-api/internal/pkg/log"
)

type queryUsecase struct {
	itemRepositoryQuery item.PostgresRepositoryQuery
	logger              log.Logger
}

func NewQueryUsecase(
	prq item.PostgresRepositoryQuery,
	log log.Logger) item.UsecaseQuery {
	return queryUsecase{
		itemRepositoryQuery: prq,
		logger:              log,
	}
}

func (q queryUsecase) GetItems(ctx context.Context) ([]response.Item, error) {
	respItem := <-q.itemRepositoryQuery.FindAllItems(ctx)
	if respItem.Error != nil {
		return nil, errors.NotFound("No items found")
	}

	itemData, ok := respItem.Data.([]entity.Item)

	if !ok {
		return nil, errors.InternalServerError("Cannot parse item data")
	}

	var resp []response.Item

	for _, item := range itemData {
		item := response.Item{
			ItemID:      item.ItemID,
			Price:       item.Price,
			Description: item.Description,
			Quantity:    item.Quantity,
		}

		resp = append(resp, item)
	}

	return resp, nil
}
