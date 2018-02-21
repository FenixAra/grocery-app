package models

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/pborman/uuid"
)

const (
	InventoryAvailable = "Available"
	InventoryBlocked   = "Blocked"
	InventorySold      = "Sold"
)

type Inventory struct {
	BarCode     string
	ItemID      string
	Name        string
	Description string
	Status      string
}

func NewInventories(i *dtos.AddInventory) []Inventory {
	var inventories []Inventory
	if i.Count > 0 {
		for ix := 0; ix < i.Count; ix++ {
			inventories = append(inventories, Inventory{
				BarCode: uuid.New(),
				ItemID:  i.ItemID,
				Status:  InventoryAvailable,
			})
		}
	} else {
		for _, inventory := range i.Inventories {
			inventories = append(inventories, Inventory{
				BarCode:     inventory.BarCode,
				ItemID:      i.ItemID,
				Status:      InventoryAvailable,
				Name:        inventory.Name,
				Description: inventory.Description,
			})
		}
	}

	return inventories
}
