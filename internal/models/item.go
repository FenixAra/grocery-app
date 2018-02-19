package models

import "github.com/FenixAra/grocery-app/dtos"

type Item struct {
	ID          string
	Name        string
	Code        string
	Description string
	PriceCardID string
	Tags        []string
}

func NewItem(i *dtos.Item) *Item {
	return &Item{
		ID:          i.ID,
		Name:        i.Name,
		Code:        i.Code,
		Description: i.Description,
		PriceCardID: i.PriceCardID,
		Tags:        i.Tags,
	}
}
