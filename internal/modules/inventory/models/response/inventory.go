package response


type InventoryData struct {
	ItemID				string				`json:"item_id"`	
	Quantity 			int64					`json:"quantity"`
	Description 	string				`json:"description"`
	Price 				int64					`json:"price"`
}