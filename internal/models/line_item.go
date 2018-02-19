package models

import "github.com/FenixAra/grocery-app/dtos"

type LineItem struct {
	ID          string
	PriceCardID string
	Name        string
	Code        string
	Description string
	Amount      int
}

func NewLineItem(li *dtos.LineItem) *LineItem {
	return &LineItem{
		ID:          li.ID,
		PriceCardID: li.PriceCardID,
		Name:        li.Name,
		Code:        li.Code,
		Description: li.Description,
		Amount:      li.Amount,
	}
}
