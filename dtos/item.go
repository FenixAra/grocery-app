package dtos

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	PriceCard   PriceCard `json:"price_card"`
	Categories  []string  `json:"categories"`
	Tags        []string  `json:"-"`
	PriceCardID string    `json:"-"`
}

type PriceCard struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Total       int        `json:"total"`
	LineItems   []LineItem `json:"line_item"`
}

type LineItem struct {
	ID          string `json:"id"`
	PriceCardID string `json:"price_card_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
