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
	INSERT INTO transactions (discount_type, discount_percent, total_price, total_discount, payment_id, customer_name, timestamp, location, payment_status, down_payment)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING transaction_id`

	timestamp := tr.Timestamp
	if timestamp.IsZero() {
		timestamp = time.Now()
	}

	err := c.postgres.QueryRow(context.Background(), query, tr.DiscountType, tr.DiscountPercent,
		tr.TotalDiscount, tr.TotalPrice, tr.PaymentID, tr.CustomerName, timestamp,
		tr.Location, tr.PaymentStatus, tr.DownPayment).Scan(&tr.TransactionID)

	if err != nil {
		return 0, err
	}

	return tr.TransactionID, nil
}

func (c commandPostgresRepository) ModifyPaymentStatus(ctx context.Context, transactionId string, paymentStatus string) error {
	query := `
	UPDATE transactions
	SET payment_status = $1
	WHERE transaction_id = $2`

	_, err := c.postgres.Query(ctx, query, paymentStatus, transactionId)

	if err != nil {
		return err
	}

	return nil
}

func (c commandPostgresRepository) RemoveTransaction(ctx context.Context, transactionId string) error {
	query := `
	DELETE FROM transactions
	WHERE transaction_id = $1`

	_, err := c.postgres.Exec(ctx, query, transactionId)
	if err != nil {
		return err
	}

	return nil
}
