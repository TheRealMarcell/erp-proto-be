package model

import "time"

type Sale struct{
	ArticleID int64								`json:"article_id"`
	TransactionNumber int64				`json:"transaction_number"`
	Description string						`json:"description"`
	Price int64										`json:"price"`
	Quantity int64								`json:"quantity"`
	Total int64										`json:"total"`
	PaymentMethod	string					`json:"payment_method"`
	Time time.Time								`json:"time"`
}



