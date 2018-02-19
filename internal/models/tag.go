package models

const (
	TypeCategory = "Category"
	TypeItem     = "Item"
)

type Tag struct {
	ID   string
	Name string
	Type string
}
