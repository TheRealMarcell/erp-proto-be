package item

import (
	"context"
	"erp-api/internal/modules/item/models/entity"
	"erp-api/internal/modules/item/models/request"
	"erp-api/internal/modules/item/models/response"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetItems(ctx context.Context) ([]response.Item, error)
}

type UsecaseCommand interface {
	SaveItem(ctx context.Context, payload request.SubmitItem) error
}

type PostgresRepositoryQuery interface {
	FindAllItems(ctx context.Context) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	InsertItems(ctx context.Context, items []entity.Item) error
}
