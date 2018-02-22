package services

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
)

type Account struct {
	l       *log.Logger
	dbConn  *db.DBConn
	account *daos.Account
}

func NewAccount(l *log.Logger, dbConn *db.DBConn) *Account {
	return &Account{
		l:       l,
		dbConn:  dbConn,
		account: daos.NewAccount(l, dbConn),
	}
}

func (a *Account) Save(account *dtos.Account) error {
	return a.account.Upsert(models.NewAccount(account))
}

func (a *Account) ChangeType(req *dtos.ChangeAccountType) error {
	account, err := a.account.Get(req.ID)
	if err != nil {
		return err
	}

	account.Type = req.Type
	return a.account.Upsert(account)
}
