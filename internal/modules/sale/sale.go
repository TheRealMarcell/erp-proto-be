package sale

import (
	"context"
	"erp-api/internal/modules/sale/models/response"

	itemRequest "erp-api/internal/modules/item/models/request"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetSales(ctx context.Context) ([]response.GetSaleResponse, error)
}

type PostgresRepositoryQuery interface {
	FindAllSales(ctx context.Context) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	BatchUpdateReturQty(ctx context.Context, items itemRequest.UpdateItem) error
}
