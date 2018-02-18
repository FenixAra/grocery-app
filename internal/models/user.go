package models

const (
	UserEmployee = "Employee"
	UserCustomer = "Customer"
)

type User struct {
	ID     string
	Name   string
	Email  string
	Mobile string
	Type   string
	Age    int
}
