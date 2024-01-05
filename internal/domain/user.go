package domain

type User struct {
	ID       string `bson:'_id'`
	Login    string `bson:'Login'`
	Password string `bson:'Password'`
}
