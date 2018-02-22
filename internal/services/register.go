package services

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
)

type Register struct {
	l        *log.Logger
	dbConn   *db.DBConn
	register *daos.Register
}

func NewRegister(l *log.Logger, dbConn *db.DBConn) *Register {
	return &Register{
		l:        l,
		dbConn:   dbConn,
		register: daos.NewRegister(l, dbConn),
	}
}

func (r *Register) Save(req *dtos.Register) error {
	registers := models.NewRegisters(req)
	for _, register := range registers {
		err := r.register.Upsert(&register)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Register) OccupyRegister(req *dtos.OccupyRegister) error {
	register, err := r.register.Get(req.ID)
	if err != nil {
		return err
	}

	register.AccountID.String = req.AccountID
	if req.AccountID != "" {
		register.AccountID.Valid = true
	}
	register.Status = req.Status
	return r.register.Upsert(register)
}
