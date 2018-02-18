package models

import "time"

type Order struct {
	ID          string
	UserID      string
	RegisterID  string
	EmployeeID  string
	Amount      int
	Bill        string
	Inventories []string
	CreatedAt   time.Time
}
