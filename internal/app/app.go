package app

import (
	"fmt"
	"todoApp/internal/config"
)

func Start() {
	cfg := config.MustLoad()

	fmt.Println(*cfg)

	// log = logger.New()
	// db = storage.New()
	// server = server.New()
	// server.Listen()
}
