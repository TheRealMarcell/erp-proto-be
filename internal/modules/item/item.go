package item

import (
	"context"
	"erp-api/internal/modules/item/models/response"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetItems(ctx context.Context) ([]response.Item, error)
}

type UsecaseCommand interface {
}

type PostgresRepositoryQuery interface {
	FindAllItems(ctx context.Context) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
}
