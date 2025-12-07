package commands

import (
	"context"
	"erp-api/internal/modules/sale"
	"erp-api/internal/modules/sale/models/entity"
	wrapper "erp-api/internal/pkg/helpers"
	"erp-api/internal/pkg/log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type queryPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewQueryPostgresRepository(
	postgres *pgxpool.Pool,
	log log.Logger) sale.PostgresRepositoryQuery {
	return queryPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (q queryPostgresRepository) FindAllSales(ctx context.Context, tid string) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	var query string

	if tid == "" {
		query = `
		SELECT sale_id, description, quantity, price, total, discount_per_item, quantity_retur, transactions.transaction_id, item_id, location
		FROM sales INNER JOIN transactions on sales.transaction_id = transactions.transaction_id
		ORDER BY sale_id
		`
	} else {
		query = `
		SELECT sale_id, description, quantity, price, total, discount_per_item, quantity_retur, transactions.transaction_id, item_id, location
		FROM sales INNER JOIN transactions on sales.transaction_id = transactions.transaction_id
		WHERE sales.transaction_id = $1
		ORDER BY sale_id
		`
	}

	var sales []entity.Sale

	go func() {
		defer close(output)

		var rows pgx.Rows
		var err error

		if tid == "" {
			rows, err = q.postgres.Query(ctx, query)
		} else {
			rows, err = q.postgres.Query(ctx, query, tid)
		}

		if err != nil {
			output <- wrapper.Result{Error: err}
			return
		}

		for rows.Next() {
			var sale entity.Sale
			err = rows.Scan(&sale.SaleID, &sale.Description, &sale.Quantity,
				&sale.Price, &sale.Total, &sale.DiscountPerItem, &sale.QuantityRetur,
				&sale.TransactionID, &sale.ItemID, &sale.Location)

			if err != nil {
				output <- wrapper.Result{Error: err}
				return
			}

			sales = append(sales, sale)
		}

		output <- wrapper.Result{Data: sales}

	}()

	return output
}
