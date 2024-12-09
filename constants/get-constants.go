package constants

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
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
		if err := godotenv.Load(".env"); err != nil {
			panic("No .env file found")
		}

		port := fmt.Sprintf(":%s", os.Getenv("PORT"))

		if port != ":" {
			port = ":3000"
		}

		constantsInstace = Constants{
			PORT: port,
		}
	})

	return constantsInstace
}
