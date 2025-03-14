package commands

import (
	"context"
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/entity"
	"erp-api/internal/modules/item/models/request"
	"erp-api/internal/pkg/log"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type commandPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewCommandPostgresRepository(postgres *pgxpool.Pool, log log.Logger) item.PostgresRepositoryCommand {
	return &commandPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (c commandPostgresRepository) BatchInsertItems(ctx context.Context, items []entity.Item) error {
	tx, err := c.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}

	for _, item := range items {
		batch.Queue("INSERT INTO items (item_id, description, price) VALUES ($1, $2, $3)",
			item.ItemID, item.Description, item.Price)
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

func (c commandPostgresRepository) ModifyItemPrice(ctx context.Context, price request.ItemPrice) error {
	query := `
	UPDATE items
	SET price = $1
	WHERE item_id = $2
	`

	_, err := c.postgres.Query(ctx, query, price.Price, price.ItemID)
	if err != nil {
		return err
	}

	return nil
}
