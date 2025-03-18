package command

import (
	"context"
	"erp-api/internal/modules/history"
	"erp-api/internal/modules/history/models/entity"
	"erp-api/internal/pkg/log"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type commandPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewCommandPostgresRepository(postgres *pgxpool.Pool, log log.Logger) history.PostgresRepositoryCommand {
	return &commandPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (c commandPostgresRepository) BatchInsertHistory(ctx context.Context, history []entity.History) error {
	tx, err := c.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	defer tx.Rollback(ctx)
	batch := &pgx.Batch{}

	groupID := uuid.New().String()

	for _, historyItem := range history {
		query := `INSERT INTO history_pindahan (item_id, quantity, timestamp, source, destination, group_id)
		VALUES ($1, $2, $3, $4, $5, $6)`
		batch.Queue(query, historyItem.ItemID, historyItem.Quantity,
			historyItem.Timestamp, historyItem.Source, historyItem.Destination, groupID)
	}

	results := tx.SendBatch(ctx, batch)

	if err := results.Close(); err != nil {
		return fmt.Errorf("batch insert error: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("transaction commit error: %v", err)
	}

	return nil
}
