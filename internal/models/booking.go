package models

import (
	"time"

	"github.com/FenixAra/grocery-app/dtos"
)

type Booking struct {
	ID          string
	AccountID   string
	RegisterID  string
	EmployeeID  string
	Amount      int
	Bill        dtos.Bill
	Inventories []string
	CreatedAt   time.Time
}
