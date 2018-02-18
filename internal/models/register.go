package models

const (
	RegisterClosed    = "RegisterClosed"
	RegisterOccupied  = "RegisterOccupied"
	RegisterAvailable = "RegisterAvailable"
)

type Register struct {
	ID     string
	UserID string
	Status string
}
