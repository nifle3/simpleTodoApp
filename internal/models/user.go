package models

type User struct {
	ID       string `json:"ID,omitempty"`
	Login    string `json:"Login,omitempty"`
	Password string `json:"Password,omitempty"`
	Email    string `json:"Email,omitemptys"`
}
