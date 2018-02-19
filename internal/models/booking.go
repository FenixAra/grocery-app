package models

import "time"

type Booking struct {
	ID          string
	AccountID   string
	RegisterID  string
	EmployeeID  string
	Amount      int
	Bill        string
	Inventories []string
	CreatedAt   time.Time
}
