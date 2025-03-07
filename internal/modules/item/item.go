package item

import (
	"context"
	wrapper "erp-api/internal/pkg/helpers"
)

type UsecaseQuery interface {
	GetAllItems(ctx context.Context)
}

type UsecaseCommand interface {

}

type PostgresRepositoryQuery interface {
	GetAllItems(ctx context.Context) <-chan wrapper.Result
}

type PostgresRepositoryCommand interface {
	
}