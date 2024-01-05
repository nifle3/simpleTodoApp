package domain

type Todo struct {
	Name        string `bson:'Name' json:'Name'`
	UserId      string
	DeadLine    string `bson:'DeadLine' json:'DeadLine'`
	Description string `bson:'Description' json:'Description'`
	IsComplete  bool   `bson:'IsComplete'`
}
