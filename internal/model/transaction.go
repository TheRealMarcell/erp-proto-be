package model

import (
	db "erp-api/database"
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
	Location     string       `json:"location"`
	PaymentStatus string			`json:"payment_status"`
}

// return type for transaction GET (no sale)
type TransactionResponse struct {
	TransactionID int64 			`json:"transaction_id"`
	DiscountType string				`json:"discount_type"`
	DiscountPercent int64			`json:"discount_percent"`
	TotalDiscount int64				`json:"total_discount"`
	PaymentID int64						`json:"payment_id"`
	CustomerName string				`json:"customer_name"`
	Timestamp time.Time				`json:"timestamp"`
	Location     string       `json:"location"`
	PaymentStatus string			`json:"payment_status"`
}

func GetTransactions() ([]TransactionResponse, error){
	query := `
	SELECT * FROM transactions
	ORDER BY transaction_id`

	rows, err := db.DB.Query(query)
	if err != nil{
		return nil, err
	}

	var transactions []TransactionResponse

	for rows.Next(){
		var tr TransactionResponse
		err = rows.Scan(&tr.TransactionID, &tr.DiscountType, &tr.DiscountPercent,
		&tr.TotalDiscount, &tr.PaymentID, &tr.CustomerName, &tr.Timestamp, &tr.Location,
		&tr.PaymentStatus)

		transactions = append(transactions, tr)
	}
	if err != nil{
		return nil, err
	}

	return transactions, nil
}

func (tr *Transaction) Save() error {
	query := `
	INSERT INTO transactions (discount_type, discount_percent, total_discount, payment_id, customer_name, timestamp, location, payment_status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING transaction_id`

	currentTime := time.Now()

	err := db.DB.QueryRow(query, tr.DiscountType, tr.DiscountPercent,
	tr.TotalDiscount, tr.PaymentID, tr.CustomerName, currentTime , tr.Location, tr.PaymentStatus).Scan(&tr.TransactionID)

	if err != nil{
		return err
	}

	// save the sale associated with the transactions
	for _, sale := range tr.Sales{
		err := sale.Save(tr.TransactionID, tr.Location)

		if err != nil{
			return err
		}
	}

	return nil
}