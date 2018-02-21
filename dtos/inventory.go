package dtos

type AddInventory struct {
	ItemID      string      `json:"item_id"`
	Count       int         `json:"count"`
	Inventories []Inventory `json:"inventories"`
}

type Inventory struct {
	BarCode     string `json:"bar_code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemID      string
}
