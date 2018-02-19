package models

const (
	UserEmployee = "Employee"
	UserCustomer = "Customer"
)

type Account struct {
	ID     string
	Name   string
	Email  string
	Mobile string
	Type   string
	Age    int
}
