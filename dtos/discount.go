package dtos

type Discount struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Amount      int      `json:"amount"`
	Percent     int      `json:"percent"`
	Inclusion   []string `json:"inclusion"`
	Exclusion   []string `json:"exclusion"`
}
