package models

type Order struct {
	ID          string
	UserID      string
	RegisterID  string
	EmployeeID  string
	Amount      int
	Bill        string
	Inventories []string
}
