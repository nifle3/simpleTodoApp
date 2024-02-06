package models

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}
