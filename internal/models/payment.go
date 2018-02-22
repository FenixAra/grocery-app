package models

type Payment struct {
	ID         string
	BookingID  string
	Mode       string
	PaymentRef string
	Amount     string
	Status     string
}
