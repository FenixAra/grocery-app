package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Discount struct {
	l  *log.Logger
	db *db.DBConn
}

func NewDiscount(l *log.Logger, db *db.DBConn) *Discount {
	return &Discount{
		l:  l,
		db: db,
	}
}

func (v *Discount) Persist(Discount *models.Discount) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Discount VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Discount.ID), qa.Append(Discount.Name), qa.Append(Discount.Code), qa.Append(Discount.Description),
		qa.Append(Discount.Type), qa.Append(Discount.Amount), qa.Append(Discount.Percent),
		qa.Append(Discount.Inclusion), qa.Append(Discount.Exclusion))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Discount) Upsert(Discount *models.Discount) error {
	err := v.Persist(Discount)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Discount SET name = %s, description = %s, code = %s, type = %s,
		amount = %s, percent = %s, inclusion = %s, exclusion = %s 
		 WHERE id = %s`, qa.Append(Discount.Name), qa.Append(Discount.Description), qa.Append(Discount.Code),
			qa.Append(Discount.Type), qa.Append(Discount.Amount), qa.Append(Discount.Percent),
			qa.Append(Discount.Inclusion), qa.Append(Discount.Exclusion), qa.Append(Discount.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Discount) Get(id string) (*models.Discount, error) {
	Discount := &models.Discount{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Discount WHERE id = $1`, id).Scan(
		&Discount.ID,
		&Discount.Name,
		&Discount.Code,
		&Discount.Description,
		&Discount.Type,
		&Discount.Amount,
		&Discount.Percent,
		&Discount.Inclusion,
		&Discount.Exclusion,
	)
	if err != nil {
		return nil, err
	}

	return Discount, nil
}

func (v *Discount) GetForTags(tags []string) ([]models.Discount, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Discount WHERE 
		(inclusion = '{}' OR $1 @> inclusion) AND NOT($1 && exclusion)`, tags)
	if err != nil {
		return nil, err
	}

	var Discounts []models.Discount
	for rows.Next() {
		var Discount models.Discount
		err = rows.Scan(
			&Discount.ID,
			&Discount.Name,
			&Discount.Code,
			&Discount.Description,
			&Discount.Type,
			&Discount.Amount,
			&Discount.Percent,
			&Discount.Inclusion,
			&Discount.Exclusion,
		)
		if err != nil {
			return nil, err
		}

		Discounts = append(Discounts, Discount)
	}

	return Discounts, nil
}

func (v *Discount) GetAll() ([]models.Discount, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Discount`)
	if err != nil {
		return nil, err
	}

	var Discounts []models.Discount
	for rows.Next() {
		var Discount models.Discount
		err = rows.Scan(
			&Discount.ID,
			&Discount.Name,
			&Discount.Code,
			&Discount.Description,
			&Discount.Type,
			&Discount.Amount,
			&Discount.Percent,
			&Discount.Inclusion,
			&Discount.Exclusion,
		)
		if err != nil {
			return nil, err
		}

		Discounts = append(Discounts, Discount)
	}

	return Discounts, nil
}
