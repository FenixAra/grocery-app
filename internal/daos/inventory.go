package daos

import (
	"errors"
	"fmt"
	"strings"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

var (
	ErrUpdateFailedForFew = errors.New("Update failed for few")
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

func (v *Inventory) InventoryStatus(barCodes []string, status, prevStatus string) error {
	qa := pgx.QueryArgs{}
	var barCodeQAs []string
	q := `UPDATE Inventory SET status = ` + qa.Append(status)
	for _, barCode := range barCodes {
		barCodeQAs = append(barCodeQAs, qa.Append(barCode))
	}
	q += fmt.Sprintf(` WHERE bar_code IN (%s) AND status = %s`, strings.Join(barCodeQAs, ","), qa.Append(prevStatus))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if int(ct.RowsAffected()) != len(barCodes) {
		return ErrUpdateFailedForFew
	}

	return nil
}

func (v *Inventory) GetAvailable(itemID string) ([]string, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT bar_code FROM Inventory WHERE item_id = $1 AND status = $2`,
		itemID, models.InventoryAvailable)
	if err != nil {
		return nil, err
	}

	var Inventories []string
	for rows.Next() {
		var Inventory string
		err = rows.Scan(
			&Inventory,
		)
		if err != nil {
			return nil, err
		}

		Inventories = append(Inventories, Inventory)
	}

	return Inventories, nil
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
