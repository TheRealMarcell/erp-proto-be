package history

import (
	"context"
	"erp-api/internal/modules/history/models/entity"
	"erp-api/internal/modules/history/models/response"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseCommand interface {
	// CURRENTLY UNUSED
	SaveHistory(ctx context.Context, history []entity.History) error
}

type UsecaseQuery interface {
	GetHistory(ctx context.Context) ([]response.History, error)
}

type PostgresRepositoryCommand interface {
	BatchInsertHistory(ctx context.Context, history []entity.History) error
}

type PostgresRepositoryQuery interface {
	GetListHistory(ctx context.Context) <-chan wrapper.Result
}
