package models

import "github.com/FenixAra/grocery-app/dtos"

const (
	ItemDiscount = "Item"
	UserDiscount = "User"
)

type Discount struct {
	ID          string
	Name        string
	Code        string
	Description string
	Type        string
	Amount      int
	Percent     int
	Inclusion   []string
	Exclusion   []string
}

func NewDiscount(d *dtos.Discount) *Discount {
	return &Discount{
		ID:          d.ID,
		Name:        d.Name,
		Code:        d.Code,
		Description: d.Description,
		Type:        d.Type,
		Amount:      d.Amount,
		Percent:     d.Percent,
		Inclusion:   d.Inclusion,
		Exclusion:   d.Exclusion,
	}
}
