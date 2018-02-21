package models

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/pborman/uuid"
	pgx "gopkg.in/jackc/pgx.v2"
)

const (
	RegisterClosed    = "RegisterClosed"
	RegisterOccupied  = "RegisterOccupied"
	RegisterAvailable = "RegisterAvailable"
)

type Register struct {
	ID        string
	AccountID pgx.NullString
	Status    string
}

func NewRegisters(r *dtos.Register) []Register {
	var registers []Register
	for i := 0; i < r.Count; i++ {
		registers = append(registers, Register{
			ID:     uuid.New(),
			Status: RegisterClosed,
			AccountID: pgx.NullString{
				String: "",
				Valid:  false,
			},
		})
	}

	return registers
}
