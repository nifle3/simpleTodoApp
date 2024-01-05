package domain

import (
	"fmt"
	"time"
)

type Todo struct {
	Name        string `bson:'Name' json:'Name'`
	UserId      string
	DeadLine    time.Time `bson:'DeadLine' json:'DeadLine'`
	Description string    `bson:'Description' json:'Description'`
	IsComplete  bool      `bson:'IsComplete'`
}

func CheckName(Name string) error {

	return nil
}

func CheckDeadLine(DeadLine time.Time) error {
	if time.Now().After(DeadLine) {
		return fmt.Errorf("Deadline is behind")
	}

	return nil
}
