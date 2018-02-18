package models

type LineItem struct {
	ID          string
	PriceCardID string
	Name        string
	Code        string
	Description string
	Amount      int
}
