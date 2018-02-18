package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type LineItem struct {
	l  *log.Logger
	db *db.DBConn
}

func NewLineItem(l *log.Logger, db *db.DBConn) *LineItem {
	return &LineItem{
		l:  l,
		db: db,
	}
}

func (v *LineItem) Persist(LineItem *models.LineItem) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Line_Item VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(LineItem.ID), qa.Append(LineItem.PriceCardID), qa.Append(LineItem.Name), qa.Append(LineItem.Code),
		qa.Append(LineItem.Description), qa.Append(LineItem.Amount))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *LineItem) Upsert(LineItem *models.LineItem) error {
	err := v.Persist(LineItem)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE LineItem SET name = %s, description = %s, code = %s, price_card_id = %s,
		amount = %s WHERE id = %s`, qa.Append(LineItem.Name), qa.Append(LineItem.Description),
			qa.Append(LineItem.Code), qa.Append(LineItem.PriceCardID), qa.Append(LineItem.Amount),
			qa.Append(LineItem.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *LineItem) Get(id string) (*models.LineItem, error) {
	LineItem := &models.LineItem{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM LineItem WHERE id = $1`, id).Scan(
		&LineItem.ID,
		&LineItem.PriceCardID,
		&LineItem.Name,
		&LineItem.Code,
		&LineItem.Description,
		&LineItem.Amount,
	)
	if err != nil {
		return nil, err
	}

	return LineItem, nil
}

func (v *LineItem) GetAll() ([]models.LineItem, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM LineItem`)
	if err != nil {
		return nil, err
	}

	var LineItems []models.LineItem
	for rows.Next() {
		var LineItem models.LineItem
		err = rows.Scan(
			&LineItem.ID,
			&LineItem.PriceCardID,
			&LineItem.Name,
			&LineItem.Code,
			&LineItem.Description,
			&LineItem.Amount,
		)
		if err != nil {
			return nil, err
		}

		LineItems = append(LineItems, LineItem)
	}

	return LineItems, nil
}
