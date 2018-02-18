package models

const (
	InventoryAvailable = "Available"
	InventoryBlocked   = "Blocked"
	InventorySold      = "Sold"
)

type Inventory struct {
	BarCode     string
	ItemID      string
	Name        string
	Descritpion string
	Status      string
}
