package commands

import (
	"context"
	"erp-api/internal/pkg/log"
	"fmt"

	itemRequest "erp-api/internal/modules/item/models/request"
	"erp-api/internal/modules/sale"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type commandPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewCommandPostgresRepository(
	postgres *pgxpool.Pool,
	log log.Logger) sale.PostgresRepositoryCommand {
	return commandPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (c commandPostgresRepository) BatchUpdateReturQty(ctx context.Context, items itemRequest.UpdateItem) error {
	tx, err := c.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}

	for _, item := range items.Items {
		query := `UPDATE sales 
		SET quantity_retur = quantity_retur + $1 
		WHERE sale_id = $2`
		batch.Queue(query, item.Quantity, item.SaleID)
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
