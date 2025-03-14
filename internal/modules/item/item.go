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
	UpdateItem(ctx context.Context, payload request.UpdateItem) error
	CorrectItem(ctx context.Context, payload request.CorrectItem, id string) error
	BrokenItem(ctx context.Context, payload request.UpdateItem) error
	UpdateItemPrice(ctx context.Context, payload request.ItemPrice) error
}

type PostgresRepositoryQuery interface {
	FindAllItems(ctx context.Context) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	BatchInsertItems(ctx context.Context, items []entity.Item) error
	ModifyItemPrice(ctx context.Context, price request.ItemPrice) error
}
