package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
}
