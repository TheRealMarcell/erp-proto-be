package commands

import (
	"context"
	"erp-api/internal/pkg/log"
	"fmt"

	itemRequest "erp-api/internal/modules/item/models/request"
	"erp-api/internal/modules/sale"
	"erp-api/internal/modules/sale/models/entity"

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

func (c commandPostgresRepository) BatchInsertSales(ctx context.Context, sales []entity.Sale, transactionId int64) error {
	tx, err := c.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}

	for _, sale := range sales {
		query := `INSERT INTO sales (item_id, description, quantity, price, total, discount_per_item, quantity_retur, transaction_id) 
		VALUES ($1, $2, $3, $4, $5, $6, 0, $7)`
		batch.Queue(query, sale.ItemID, sale.Description,
			sale.Quantity, sale.Price, sale.Total, sale.DiscountPerItem, transactionId)
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

func (c commandPostgresRepository) BatchDeleteSales(ctx context.Context, sales []entity.Sale) error {
	tx, err := c.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}

	for _, sale := range sales {
		query := `DELETE FROM sales
		WHERE sale_id = $1
		`
		batch.Queue(query, sale.SaleID)
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
