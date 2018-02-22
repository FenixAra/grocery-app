package daos

import (
	"errors"
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

var (
	ErrNoRowsUpdated = errors.New("No Rows Updated")
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
		qa.Append(Register.ID), qa.Append(Register.AccountID), qa.Append(Register.Status))
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
		q := `UPDATE Register SET `
		q += `account_id = ` + qa.Append(Register.AccountID) + `,`
		q += fmt.Sprintf(` status = %s 
		 WHERE id = %s`, qa.Append(Register.Status), qa.Append(Register.ID))
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
		&Register.AccountID,
		&Register.Status,
	)
	if err != nil {
		return nil, err
	}

	return Register, nil
}

func (v *Register) SetStatus(id, status, prevStatus string) error {
	ct, err := v.db.GetQueryer().Exec(`UPDATE register SET status = $1 WHERE id = $2 AND status = $3`, status, id, prevStatus)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsUpdated
	}

	return nil
}

func (v *Register) GetAvailable() ([]string, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT id FROM Register WHERE status = $1`, models.RegisterAvailable)
	if err != nil {
		return nil, err
	}

	var Registers []string
	for rows.Next() {
		var Register string
		err = rows.Scan(
			&Register,
		)
		if err != nil {
			return nil, err
		}

		Registers = append(Registers, Register)
	}

	return Registers, nil
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
			&Register.AccountID,
			&Register.Status,
		)
		if err != nil {
			return nil, err
		}

		Registers = append(Registers, Register)
	}

	return Registers, nil
}
