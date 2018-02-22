package dtos

type BookingData struct {
	Items     map[string]int `json:"items"`
	AccountID string         `json:"account_id"`
}

type GenerateBillResponse struct {
	RegisterID  string   `json:"register_id"`
	BookingID   string   `json:"booking_id"`
	AccountID   string   `json:"account_id"`
	Amount      int      `json:"amount"`
	Inventories []string `json:"inventories"`
	Bill        Bill     `json:"bill"`
}

type Bill struct {
	Amount    int            `json:"amount"`
	LineItems []BillLineItem `json:"breakup"`
	Discounts []BillDiscount `json:"discounts"`
}

type BillLineItem struct {
	ID          string `json:"id"`
	PriceCardID string `json:"price_card_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type BillDiscount struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
