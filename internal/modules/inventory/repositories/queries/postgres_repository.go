package queries

import (
	"context"
	"erp-api/internal/modules/inventory"
	"erp-api/internal/modules/inventory/models/entity"
	wrapper "erp-api/internal/pkg/helpers"
	"erp-api/internal/pkg/log"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type queryPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewQueryPostgresRepository(postgres *pgxpool.Pool, log log.Logger) inventory.PostgresRepositoryQuery {
	return &queryPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (c queryPostgresRepository) FindListInventory(ctx context.Context, location string) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	go func() {
		defer close(output)

		// sanitise and format location
		if err := wrapper.IsValidLocation(location); err != nil {
			output <- wrapper.Result{Error: err}
		}

		formatted_location := fmt.Sprintf("inventory_%v", location)

		query := fmt.Sprintf(`
		SELECT i.item_id, quantity, i.description, price 
		FROM %s i
		INNER JOIN items ON i.item_id = items.item_id
		`, formatted_location)

		rows, err := c.postgres.Query(ctx, query)
		if err != nil {
			output <- wrapper.Result{Error: err}
			return
		}

		var inventoryItems []entity.InventoryItem
		for rows.Next() {
			var item entity.InventoryItem
			err = rows.Scan(&item.ItemID, &item.Quantity, &item.Description, &item.Price)
			if err != nil {
				output <- wrapper.Result{Error: err}
				return
			}
			inventoryItems = append(inventoryItems, item)
		}

		output <- wrapper.Result{Data: inventoryItems}
	}()

	return output
}
