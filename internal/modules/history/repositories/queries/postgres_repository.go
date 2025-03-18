package query

import (
	"context"
	"erp-api/internal/modules/history"
	"erp-api/internal/modules/history/models/entity"
	"erp-api/internal/pkg/log"

	wrapper "erp-api/internal/pkg/helpers"

	"github.com/jackc/pgx/v5/pgxpool"
)

type queryPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewQueryPostgresRepository(postgres *pgxpool.Pool, log log.Logger) history.PostgresRepositoryQuery {
	return &queryPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (q queryPostgresRepository) GetListHistory(ctx context.Context) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		defer close(output)

		query := `SELECT *
		FROM history_pindahan`

		rows, err := q.postgres.Query(ctx, query)
		if err != nil {
			output <- wrapper.Result{Error: err}
			return
		}

		var historyItems []entity.History
		for rows.Next() {
			var historyItem entity.History
			err = rows.Scan(&historyItem.PindahanID, &historyItem.ItemID, &historyItem.Quantity,
				&historyItem.Timestamp, &historyItem.Source, &historyItem.Destination, &historyItem.GroupID)
			if err != nil {
				output <- wrapper.Result{Error: err}
				return
			}
			historyItems = append(historyItems, historyItem)
		}

		output <- wrapper.Result{Data: historyItems}
	}()

	return output
}
