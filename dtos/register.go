package dtos

type Register struct {
	Count int `json:"count"`
}

type OccupyRegister struct {
	ID        string `json:"id"`
	AccountID string `json:"accountID"`
	Status    string `json:"status"`
}
