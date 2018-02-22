package models

import "github.com/FenixAra/grocery-app/dtos"

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

func NewAccount(a *dtos.Account) *Account {
	return &Account{
		ID:     a.ID,
		Name:   a.Name,
		Email:  a.Email,
		Mobile: a.Mobile,
		Age:    a.Age,
		Type:   UserCustomer,
	}
}
