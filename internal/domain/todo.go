package domain

import (
	"time"
)

type Todo struct {
	ID          string    `json:"Id"`
	Name        string    `json:"Name"`
	DeadLine    time.Time `json:"DeadLine"`
	Description string    `json:"Description"`
	IsComplete  bool      `json:"IsComplete"`
}
