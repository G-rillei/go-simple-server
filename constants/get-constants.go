package constants

import (
	"os"
	"sync"
)

type Constants struct {
	PORT string
}

var (
	constantsInstace Constants
	once             sync.Once
)

func GetConstants() Constants {
	once.Do(func() {
		port := os.Getenv("PORT")

		if port == "" {
			port = ":3000"
		} else {
			port = ":" + port
		}

		constantsInstace = Constants{
			PORT: port,
		}
	})

	return constantsInstace
}
