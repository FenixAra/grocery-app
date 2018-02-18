package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Item struct {
	l  *log.Logger
	db *db.DBConn
}

func NewItem(l *log.Logger, db *db.DBConn) *Item {
	return &Item{
		l:  l,
		db: db,
	}
}

func (v *Item) Persist(Item *models.Item) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Item VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Item.ID), qa.Append(Item.Name), qa.Append(Item.Code), qa.Append(Item.Description),
		qa.Append(Item.PriceCardID), qa.Append(Item.Tags))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Item) Upsert(Item *models.Item) error {
	err := v.Persist(Item)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Item SET name = %s, description = %s, code = %s, price_card_id = %s,
		tags = %s   
		 WHERE id = %s`, qa.Append(Item.Name), qa.Append(Item.Description), qa.Append(Item.Code),
			qa.Append(Item.PriceCardID), qa.Append(Item.Tags), qa.Append(Item.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Item) Get(id string) (*models.Item, error) {
	Item := &models.Item{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Item WHERE id = $1`, id).Scan(
		&Item.ID,
		&Item.Name,
		&Item.Code,
		&Item.Description,
		&Item.PriceCardID,
		&Item.Tags,
	)
	if err != nil {
		return nil, err
	}

	return Item, nil
}

func (v *Item) GetAll() ([]models.Item, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Item`)
	if err != nil {
		return nil, err
	}

	var Items []models.Item
	for rows.Next() {
		var Item models.Item
		err = rows.Scan(
			&Item.ID,
			&Item.Name,
			&Item.Code,
			&Item.Description,
			&Item.PriceCardID,
			&Item.Tags,
		)
		if err != nil {
			return nil, err
		}

		Items = append(Items, Item)
	}

	return Items, nil
}
