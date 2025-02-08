package model

import (
	"context"
	db "erp-api/database"
	"fmt"
	"time"
)

type Transaction struct {
	TransactionID int64 			`json:"transaction_id"`
	Sales []Sale							`json:"sales"`
	DiscountType string				`json:"discount_type"`
	DiscountPercent int64			`json:"discount_percent"`
	TotalDiscount int64				`json:"total_discount"`
	PaymentID int64						`json:"payment_id"`
	CustomerName string				`json:"customer_name"`
	Timestamp time.Time				`json:"timestamp"`
}


func (tr *Transaction) Save() error {
	query := `
	INSERT INTO transactions (transaction_id, discount_type, discount_percent, total_discount, payment_id, customer_name, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	currentTime := time.Now()

	_, err := db.Conn.Exec(context.Background(), query, tr.TransactionID, tr.DiscountType, tr.DiscountPercent,
	tr.TotalDiscount, tr.PaymentID, tr.CustomerName, currentTime)

	if err != nil{
		return err
	}

	fmt.Println("TransactionID:", tr.TransactionID)

	// save sale corresponding to transaction id
	for _, sale := range tr.Sales{
		err := sale.Save(tr.TransactionID)

		if err != nil{
			return err
		}
	}

	return nil
}