package dtos

type Account struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Type   string `json:"-"`
	Age    int    `json:"age"`
}

type ChangeAccountType struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
