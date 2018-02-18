package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Register struct {
	l  *log.Logger
	db *db.DBConn
}

func NewRegister(l *log.Logger, db *db.DBConn) *Register {
	return &Register{
		l:  l,
		db: db,
	}
}

func (v *Register) Persist(Register *models.Register) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Register VALUES (%s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Register.ID), qa.Append(Register.UserID), qa.Append(Register.Status))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Register) Upsert(Register *models.Register) error {
	err := v.Persist(Register)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Register SET user_id = %s, status = %s 
		 WHERE id = %s`, qa.Append(Register.UserID), qa.Append(Register.Status), qa.Append(Register.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Register) Get(id string) (*models.Register, error) {
	Register := &models.Register{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Register WHERE id = $1`, id).Scan(
		&Register.ID,
		&Register.UserID,
		&Register.Status,
	)
	if err != nil {
		return nil, err
	}

	return Register, nil
}

func (v *Register) GetAll() ([]models.Register, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Register`)
	if err != nil {
		return nil, err
	}

	var Registers []models.Register
	for rows.Next() {
		var Register models.Register
		err = rows.Scan(
			&Register.ID,
			&Register.UserID,
			&Register.Status,
		)
		if err != nil {
			return nil, err
		}

		Registers = append(Registers, Register)
	}

	return Registers, nil
}
