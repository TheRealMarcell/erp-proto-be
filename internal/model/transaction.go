package model

import (
	"context"
	db "erp-api/database"
	"time"
)

// return type for transaction GET (no sale)
type TransactionResponse struct {
	TransactionID   int64     `json:"transaction_id"`
	DiscountType    string    `json:"discount_type"`
	DiscountPercent int64     `json:"discount_percent"`
	TotalDiscount   int64     `json:"total_discount"`
	TotalPrice      int64     `json:"total_price"`
	PaymentID       int64     `json:"payment_id"`
	CustomerName    string    `json:"customer_name"`
	Timestamp       time.Time `json:"timestamp"`
	Location        string    `json:"location"`
	PaymentStatus   string    `json:"payment_status"`
}

type DiscountResponse struct {
	TransactionID   int64 `json:"transaction_id"`
	DiscountPercent int64 `json:"discount_percent"`
}

func UpdatePaymentStatus(transaction_id string, payment_status string) error {
	query := `
	UPDATE transactions
	SET payment_status = $1
	WHERE transaction_id = $2`

	_, err := db.DB.Query(context.Background(), query, payment_status, transaction_id)
	if err != nil {
		return err
	}

	return nil
}
