package request

import (
	saleEntity "erp-api/internal/modules/sale/models/entity"
	"time"
)

type Transaction struct {
	TransactionID   int64             `json:"transaction_id"`
	Sales           []saleEntity.Sale `json:"sales"`
	DiscountType    string            `json:"discount_type"`
	DiscountPercent int64             `json:"discount_percent"`
	TotalDiscount   int64             `json:"total_discount"`
	TotalPrice      int64             `json:"total_price"`
	PaymentID       int64             `json:"payment_id"`
	CustomerName    string            `json:"customer_name"`
	Timestamp       time.Time         `json:"timestamp"`
	Location        string            `json:"location"`
	PaymentStatus   string            `json:"payment_status"`
}
