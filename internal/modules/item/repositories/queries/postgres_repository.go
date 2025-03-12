package queries

import (
	"context"
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/entity"
	wrapper "erp-api/internal/pkg/helpers"
	"erp-api/internal/pkg/log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type queryPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewQueryPostgresRepository(postgres *pgxpool.Pool, log log.Logger) item.PostgresRepositoryQuery {
	return &queryPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (q queryPostgresRepository) FindAllItems(ctx context.Context) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	var items []entity.Item

	query := `
		SELECT * 
		FROM items
	`

	go func() {
		defer close(output)

		rows, err := q.postgres.Query(context.Background(), query)
		if err != nil {
			output <- wrapper.Result{Error: err}
			return
		}

		for rows.Next() {
			var item entity.Item
			err = rows.Scan(&item.ItemID, &item.Price, &item.Description)
			if err != nil {
				output <- wrapper.Result{Error: err}
				return
			}
			items = append(items, item)
		}

		output <- wrapper.Result{Data: items}
	}()

	return output
}
