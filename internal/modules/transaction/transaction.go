package transaction

import (
	"context"
	"erp-api/internal/modules/transaction/models/request"
	"erp-api/internal/modules/transaction/models/response"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetTransactions(ctx context.Context) ([]response.Transaction, error)
	GetDiscountPercentages(ctx context.Context) ([]response.GetDiscount, error)
}

type UsecaseCommand interface {
	InsertTransaction(ctx context.Context, payload request.Transaction) error
}

type PostgresRepositoryQuery interface {
	FindAllTransactions(ctx context.Context) <-chan wrapper.Result
	FindDiscount(ctx context.Context) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	SaveTransaction(ctx context.Context, transaction request.Transaction) (int64, error)
}
