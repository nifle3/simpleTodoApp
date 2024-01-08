package domain

type User struct {
	ID       string `json:'ID'`
	Login    string `json:'Login'`
	Password string `json:'Password'`
}
