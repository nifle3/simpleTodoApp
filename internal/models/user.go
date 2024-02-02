package models

type User struct {
	ID       string
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
