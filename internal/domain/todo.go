package domain

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          string    `json:"Id"`
	Name        string    `json:"Name"`
	DeadLine    time.Time `json:"DeadLine"`
	Description string    `json:"Description"`
	IsComplete  bool      `json:"IsComplete"`
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
