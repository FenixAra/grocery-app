package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Inventory struct {
	l  *log.Logger
	db *db.DBConn
}

func NewInventory(l *log.Logger, db *db.DBConn) *Inventory {
	return &Inventory{
		l:  l,
		db: db,
	}
}

func (v *Inventory) Persist(Inventory *models.Inventory) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Inventory VALUES (%s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Inventory.BarCode), qa.Append(Inventory.ItemID), qa.Append(Inventory.Name),
		qa.Append(Inventory.Description), qa.Append(Inventory.Status))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Inventory) Upsert(Inventory *models.Inventory) error {
	err := v.Persist(Inventory)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Inventory SET item_id = %s, name = %s, description = %s, status = %s 
		 WHERE bar_code = %s`, qa.Append(Inventory.ItemID), qa.Append(Inventory.Name), qa.Append(Inventory.Description),
			qa.Append(Inventory.Status), qa.Append(Inventory.BarCode))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Inventory) Get(barCode string) (*models.Inventory, error) {
	Inventory := &models.Inventory{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Inventory WHERE bar_code = $1`, barCode).Scan(
		&Inventory.BarCode,
		&Inventory.ItemID,
		&Inventory.Name,
		&Inventory.Description,
		&Inventory.Status,
	)
	if err != nil {
		return nil, err
	}

	return Inventory, nil
}

func (v *Inventory) GetAll() ([]models.Inventory, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Inventory`)
	if err != nil {
		return nil, err
	}

	var Inventories []models.Inventory
	for rows.Next() {
		var Inventory models.Inventory
		err = rows.Scan(
			&Inventory.BarCode,
			&Inventory.ItemID,
			&Inventory.Name,
			&Inventory.Description,
			&Inventory.Status,
		)
		if err != nil {
			return nil, err
		}

		Inventories = append(Inventories, Inventory)
	}

	return Inventories, nil
}
