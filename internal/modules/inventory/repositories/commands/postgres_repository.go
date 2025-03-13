package commands

import (
	"context"
	"erp-api/internal/modules/inventory"
	itemEntity "erp-api/internal/modules/item/models/entity"
	"erp-api/internal/pkg/log"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type commandPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewCommandPostgresRepository(postgres *pgxpool.Pool, log log.Logger) inventory.PostgresRepositoryCommand {
	return &commandPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (c commandPostgresRepository) BatchInsertInventory(ctx context.Context, items []itemEntity.Item) error {
	tx, err := c.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction error: %v", err)
	}

	defer tx.Rollback(ctx)
	storageBatch := &pgx.Batch{}

	for _, item := range items {
		var storageItem itemEntity.StorageItem
		storageItem.ItemID = item.ItemID
		storageItem.Description = item.Description
		storageItem.Quantity = item.Quantity

		storageBatch.Queue("INSERT INTO inventory_gudang (item_id, quantity, description) VALUES ($1, $2, $3)",
			storageItem.ItemID, storageItem.Quantity, storageItem.Description)

		// insert 0 for the other inventories for placeholder
		storageBatch.Queue("INSERT INTO inventory_tiktok (item_id, quantity, description) VALUES ($1, $2, $3)",
			storageItem.ItemID, 0, storageItem.Description)

		storageBatch.Queue("INSERT INTO inventory_toko (item_id, quantity, description) VALUES ($1, $2, $3)",
			storageItem.ItemID, 0, storageItem.Description)

		storageBatch.Queue("INSERT INTO inventory_rusak (item_id, quantity, description) VALUES ($1, $2, $3)",
			storageItem.ItemID, 0, storageItem.Description)
	}

	results := tx.SendBatch(ctx, storageBatch)

	if err := results.Close(); err != nil {
		return fmt.Errorf("batch insert error: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("transaction commit error: %v", err)
	}

	return nil
}
