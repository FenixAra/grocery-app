package models

type Item struct {
	ID          string
	Name        string
	Code        string
	Description string
	PriceCardID string
	Tags        []string
}
