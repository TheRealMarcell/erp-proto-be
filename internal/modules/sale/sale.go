package sale

import (
	"context"
	"erp-api/internal/modules/sale/models/response"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetSales(ctx context.Context) ([]response.GetSaleResponse, error)
}

type PostgresRepositoryQuery interface {
	FindAllSales(ctx context.Context) <-chan wrapper.Result
}

