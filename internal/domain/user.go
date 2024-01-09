package domain

type User struct {
	ID       string `json:"ID,omitempty"`
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
