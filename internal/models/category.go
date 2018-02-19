package models

import "github.com/FenixAra/grocery-app/dtos"

type Category struct {
	ID          string
	Name        string
	Description string
	ParentID    string
}

func NewCategory(c *dtos.Category) *Category {
	return &Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		ParentID:    c.ParentID,
	}
}
