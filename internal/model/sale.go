package model

import (
	"context"
	db "erp-api/database"
	"time"
)

type Sale struct{
	SaleID int64									`json:"sale_id"`
	ItemID int64									`json:"item_id"`
	TransactionNumber int64				`json:"transaction_number"`
	Description string						`json:"description"`
	Quantity int64								`json:"quantity"`
	Price int64										`json:"price"`
	Total int64										`json:"total"`
	Discount float64							`json:"discount"`
	PaymentID int64								`json:"payment_id"`
	Timestamp time.Time						`json:"timestamp"`
}

func GetSales() ([] Sale, error){
	query := `
	SELECT * FROM SALES
	ORDER BY sale_id 
	`

	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil{
		return nil, err
	}

	var sales []Sale

	for rows.Next(){
		var sale Sale
		err = rows.Scan(&sale.SaleID, &sale.ItemID, &sale.TransactionNumber, &sale.Description, &sale.Quantity,
			&sale.Price, &sale.Total, &sale.Discount, &sale.PaymentID, &sale.Timestamp)
		
		sales = append(sales, sale)
	}

	if err != nil{
		return nil, err
	}


	return sales, nil
}

func (sale *Sale) Save() error{
	query := `INSERT INTO sales 
	(item_id, transaction_number, description, quantity, price, total, discount, payment_id, timestamp)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	currentTime := time.Now()

	_, err := db.Conn.Exec(context.Background(), query, sale.ItemID, sale.TransactionNumber, sale.Description, 
	sale.Quantity, sale.Price, sale.Total, sale.Discount, sale.PaymentID, currentTime)

	if err != nil{
		return err
	}

	return nil

}

