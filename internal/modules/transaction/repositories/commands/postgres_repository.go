package commands

import (
	"context"
	"erp-api/internal/modules/transaction"
	"erp-api/internal/modules/transaction/models/request"
	"erp-api/internal/pkg/log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type commandPostgresRepository struct {
	postgres *pgxpool.Pool
	logger   log.Logger
}

func NewCommandPostgresRepository(postgres *pgxpool.Pool, log log.Logger) transaction.PostgresRepositoryCommand {
	return &commandPostgresRepository{
		postgres: postgres,
		logger:   log,
	}
}

func (c commandPostgresRepository) SaveTransaction(ctx context.Context, tr request.Transaction) (int64, error) {
	query := `
	INSERT INTO transactions (discount_type, discount_percent, total_price, total_discount, payment_id, customer_name, timestamp, location, payment_status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING transaction_id`

	currentTime := time.Now()

	err := c.postgres.QueryRow(context.Background(), query, tr.DiscountType, tr.DiscountPercent,
		tr.TotalDiscount, tr.TotalPrice, tr.PaymentID, tr.CustomerName, currentTime, tr.Location, tr.PaymentStatus).Scan(&tr.TransactionID)

	if err != nil {
		return 0, err
	}

	return tr.TransactionID, nil
}
