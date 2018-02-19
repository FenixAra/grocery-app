package services

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
)

type Discount struct {
	l        *log.Logger
	dbConn   *db.DBConn
	discount *daos.Discount
}

func NewDiscount(l *log.Logger, dbConn *db.DBConn) *Discount {
	return &Discount{
		l:        l,
		dbConn:   dbConn,
		discount: daos.NewDiscount(l, dbConn),
	}
}

func (d *Discount) Save(Discount *dtos.Discount) error {
	return d.discount.Upsert(models.NewDiscount(Discount))
}
