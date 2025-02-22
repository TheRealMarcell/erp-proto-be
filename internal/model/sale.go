package model

import (
	db "erp-api/database"
	"fmt"
)

type Sale struct{
	SaleID int64									`json:"sale_id"`
	ItemID string									`json:"item_id"`
	Description string						`json:"description"`
	Quantity int64								`json:"quantity"`
	Price int64										`json:"price"`
	Total int64										`json:"total"`
	DiscountPerItem float64				`json:"discount_per_item"`
	TransactionID int64						`json:"transaction_id"`
	Location string								`json:"location"`
}

func GetSales() ([] Sale, error){
	query := `
	SELECT sale_id, description, quantity, price, total, discount_per_item, transactions.transaction_id, item_id, location
	FROM sales INNER JOIN transactions on sales.transaction_id = transactions.transaction_id
	ORDER BY sale_id
	`

	rows, err := db.DB.Query(query)
	if err != nil{
		return nil, err
	}

	var sales []Sale

	for rows.Next(){
		var sale Sale
		err = rows.Scan(&sale.SaleID, &sale.Description, &sale.Quantity,
			&sale.Price, &sale.Total, &sale.DiscountPerItem, &sale.TransactionID, &sale.ItemID, &sale.Location)
		
		sales = append(sales, sale)
	}

	if err != nil{
		return nil, err
	}

	return sales, nil
}

func (sale *Sale) Save(transaction_id int64, location string) error{
	fmt.Println(transaction_id)
	sale.TransactionID = transaction_id

	err := UpdateInventory(sale.ItemID, sale.Quantity, location)

	if err != nil {
		return err
	}

	query := `INSERT INTO sales 
	(item_id, description, quantity, price, total, discount_per_item, quantity_retur, transaction_id)
	VALUES ($1, $2, $3, $4, $5, $6, 0, $7)`

	_, err = db.DB.Exec(query, sale.ItemID, sale.Description, 
	sale.Quantity, sale.Price, sale.Total, sale.DiscountPerItem, sale.TransactionID)

	if err != nil{
		return err
	}

	return nil
}


// updates stock for the specific location
func UpdateInventory(itemID string, quantity int64, location string) error {
	fmt.Println("Updating inventory for location:", location)

	var inventoryTable string
	switch location {
	case "tiktok":
		inventoryTable = "inventory_tiktok"
	case "toko":
		inventoryTable = "inventory_toko"
	case "gudang":
		inventoryTable = "inventory_gudang"
	default:
		return fmt.Errorf("invalid location: %s", location)
	}

	// check if quantity in inventory is sufficient
	var stock int64

	query := fmt.Sprintf(`
		SELECT quantity
		FROM %s
		WHERE item_id = $1
		`, inventoryTable)

		err := db.DB.QueryRow(query, itemID).Scan(&stock)

		if err != nil{
			return err
		}

		if stock < quantity{
			return fmt.Errorf("not enough stock to process transaction")
		}

		

	// Update stock and ensure at least one row was affected
	query = fmt.Sprintf(`
		UPDATE %s 
		SET quantity = quantity - $1 
		WHERE item_id = $2 AND quantity >= $1`, inventoryTable)

	_, err = db.DB.Exec(query, quantity, itemID)
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}

	// rowsAffected := res.RowsAffected()
	// if rowsAffected == 0 {
	// 	return fmt.Errorf("insufficient stock for item: %s", itemID)
	// }

	fmt.Printf("Stock updated for %s in %s", itemID, inventoryTable)
	return nil
}
