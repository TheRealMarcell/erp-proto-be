package response

type GetSaleResponse struct {
	SaleID          int64   `json:"sale_id"`
	ItemID          string  `json:"item_id"`
	Description     string  `json:"description"`
	Quantity        int64   `json:"quantity"`
	Price           int64   `json:"price"`
	Total           int64   `json:"total"`
	DiscountPerItem float64 `json:"discount_per_item"`
	QuantityRetur   int64   `json:"quantity_retur"`
	TransactionID   int64   `json:"transaction_id"`
	Location        string  `json:"location"`
}
