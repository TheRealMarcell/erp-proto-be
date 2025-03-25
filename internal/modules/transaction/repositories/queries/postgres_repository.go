package queries

import (
	"context"
	"erp-api/internal/modules/transaction"
	"erp-api/internal/modules/transaction/models/entity"
	wrapper "erp-api/internal/pkg/helpers"
	"erp-api/internal/pkg/log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type queryPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewQueryPostgresRepository(postgres *pgxpool.Pool, log log.Logger) transaction.PostgresRepositoryQuery {
	return &queryPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (p queryPostgresRepository) FindAllTransactions(ctx context.Context) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	query := `
	SELECT * FROM transactions
	ORDER BY transaction_id`

	go func() {
		defer close(output)

		rows, err := p.postgres.Query(ctx, query)
		if err != nil {
			output <- wrapper.Result{Error: err}
			return
		}

		var transactions []entity.Transaction

		for rows.Next() {
			var tr entity.Transaction
			err = rows.Scan(&tr.TransactionID, &tr.DiscountType, &tr.DiscountPercent, &tr.TotalPrice,
				&tr.TotalDiscount, &tr.PaymentID, &tr.CustomerName, &tr.Timestamp, &tr.Location,
				&tr.PaymentStatus, &tr.PaymentStatus)

			if err != nil {
				output <- wrapper.Result{Error: err}
				return
			}

			transactions = append(transactions, tr)
		}

		output <- wrapper.Result{Data: transactions}
	}()

	return output
}

func (p queryPostgresRepository) FindDiscount(ctx context.Context) <-chan wrapper.Result {
	output := make(chan wrapper.Result)

	var discounts []entity.TransactionDiscount

	query := `
	SELECT transaction_id, discount_percent
	FROM transactions`

	go func() {
		defer close(output)

		rows, err := p.postgres.Query(ctx, query)
		if err != nil {
			output <- wrapper.Result{Error: err}
			return
		}

		for rows.Next() {
			var discount entity.TransactionDiscount
			err = rows.Scan(&discount.TransactionID, &discount.DiscountPercent)
			if err != nil {
				output <- wrapper.Result{Error: err}
				return
			}
			discounts = append(discounts, discount)
		}

		output <- wrapper.Result{Data: discounts}
	}()

	return output
}
