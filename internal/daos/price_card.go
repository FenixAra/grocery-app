package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type PriceCard struct {
	l  *log.Logger
	db *db.DBConn
}

func NewPriceCard(l *log.Logger, db *db.DBConn) *PriceCard {
	return &PriceCard{
		l:  l,
		db: db,
	}
}

func (v *PriceCard) Persist(PriceCard *models.PriceCard) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Price_Card VALUES (%s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(PriceCard.ID), qa.Append(PriceCard.Code), qa.Append(PriceCard.Name), qa.Append(PriceCard.Description),
		qa.Append(PriceCard.Total))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *PriceCard) Upsert(PriceCard *models.PriceCard) error {
	err := v.Persist(PriceCard)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Price_Card SET name = %s, description = %s, code = %s, total = %s
		 WHERE id = %s`, qa.Append(PriceCard.Name), qa.Append(PriceCard.Description), qa.Append(PriceCard.Code),
			qa.Append(PriceCard.Total), qa.Append(PriceCard.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *PriceCard) Get(id string) (*models.PriceCard, error) {
	PriceCard := &models.PriceCard{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Price_Card WHERE id = $1`, id).Scan(
		&PriceCard.ID,
		&PriceCard.Code,
		&PriceCard.Name,
		&PriceCard.Description,
		&PriceCard.Total,
	)
	if err != nil {
		return nil, err
	}

	return PriceCard, nil
}

func (v *PriceCard) GetAll() ([]models.PriceCard, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Price_Card`)
	if err != nil {
		return nil, err
	}

	var PriceCards []models.PriceCard
	for rows.Next() {
		var PriceCard models.PriceCard
		err = rows.Scan(
			&PriceCard.ID,
			&PriceCard.Code,
			&PriceCard.Name,
			&PriceCard.Description,
			&PriceCard.Total,
		)
		if err != nil {
			return nil, err
		}

		PriceCards = append(PriceCards, PriceCard)
	}

	return PriceCards, nil
}
