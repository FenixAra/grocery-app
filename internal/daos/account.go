package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Account struct {
	l  *log.Logger
	db *db.DBConn
}

func NewAccount(l *log.Logger, db *db.DBConn) *Account {
	return &Account{
		l:  l,
		db: db,
	}
}

func (v *Account) Persist(Account *models.Account) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Account VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Account.ID), qa.Append(Account.Name), qa.Append(Account.Email), qa.Append(Account.Mobile),
		qa.Append(Account.Type), qa.Append(Account.Age))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Account) Upsert(Account *models.Account) error {
	err := v.Persist(Account)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Account SET name = %s, email = %s, mobile = %s, type = %s,
		age = %s   
		 WHERE id = %s`, qa.Append(Account.Name), qa.Append(Account.Email), qa.Append(Account.Mobile),
			qa.Append(Account.Type), qa.Append(Account.Age), qa.Append(Account.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Account) Get(id string) (*models.Account, error) {
	Account := &models.Account{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Account WHERE id = $1`, id).Scan(
		&Account.ID,
		&Account.Name,
		&Account.Email,
		&Account.Mobile,
		&Account.Type,
		&Account.Age,
	)
	if err != nil {
		return nil, err
	}

	return Account, nil
}

func (v *Account) GetAll() ([]models.Account, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Account`)
	if err != nil {
		return nil, err
	}

	var Accounts []models.Account
	for rows.Next() {
		var Account models.Account
		err = rows.Scan(
			&Account.ID,
			&Account.Name,
			&Account.Email,
			&Account.Mobile,
			&Account.Type,
			&Account.Age,
		)
		if err != nil {
			return nil, err
		}

		Accounts = append(Accounts, Account)
	}

	return Accounts, nil
}
