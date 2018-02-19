package models

import "github.com/FenixAra/grocery-app/dtos"

type PriceCard struct {
	ID          string
	Code        string
	Name        string
	Description string
	Total       int
}

func NewPriceCard(pc *dtos.PriceCard) *PriceCard {
	return &PriceCard{
		ID:          pc.ID,
		Code:        pc.Code,
		Name:        pc.Name,
		Description: pc.Description,
		Total:       pc.Total,
	}
}
