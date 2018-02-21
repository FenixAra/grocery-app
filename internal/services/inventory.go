package services

import (
	"errors"

	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

var (
	ErrItemNotFound = errors.New("Item not found")
)

type Inventory struct {
	l         *log.Logger
	dbConn    *db.DBConn
	inventory *daos.Inventory
	item      *daos.Item
}

func NewInventory(l *log.Logger, dbConn *db.DBConn) *Inventory {
	return &Inventory{
		l:         l,
		dbConn:    dbConn,
		inventory: daos.NewInventory(l, dbConn),
		item:      daos.NewItem(l, dbConn),
	}
}

func (i *Inventory) Save(req *dtos.AddInventory) error {
	item, err := i.item.Get(req.ItemID)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if err == pgx.ErrNoRows {
		return ErrItemNotFound
	}

	inventories := models.NewInventories(req)
	for _, inventory := range inventories {
		if inventory.Name == "" {
			inventory.Name = item.Name
		}

		if inventory.Description == "" {
			inventory.Description = item.Description
		}

		err = i.inventory.Upsert(&inventory)
		if err != nil {
			return err
		}
	}
	return nil
}
