package models

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
