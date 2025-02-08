package model

import (
	"context"
	db "erp-api/database"
	"fmt"
)

type Sale struct{
	SaleID int64									`json:"sale_id"`
	ItemID int64									`json:"item_id"`
	Description string						`json:"description"`
	Quantity int64								`json:"quantity"`
	Price int64										`json:"price"`
	Total int64										`json:"total"`
	Discount float64							`json:"discount"`
	TransactionID int64						`json:"transaction_id"`
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
		err = rows.Scan(&sale.SaleID, &sale.ItemID,&sale.Description, &sale.Quantity,
			&sale.Price, &sale.Total, &sale.Discount)
		
		sales = append(sales, sale)
	}

	if err != nil{
		return nil, err
	}


	return sales, nil
}

func (sale *Sale) Save(transaction_id int64) error{
	fmt.Println(transaction_id)
	sale.TransactionID = transaction_id

	query := `INSERT INTO sales 
	(item_id, description, quantity, price, total, discount, transaction_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.Conn.Exec(context.Background(), query, sale.ItemID, sale.Description, 
	sale.Quantity, sale.Price, sale.Total, sale.Discount, sale.TransactionID)

	if err != nil{
		return err
	}

	return nil

}

