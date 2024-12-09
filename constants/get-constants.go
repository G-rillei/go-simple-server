package constants

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type ServerConstants struct {
	Port string
}

type DatabaseConstants struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Constants struct {
	Server   ServerConstants
	Database DatabaseConstants
}

var (
	constantsInstace Constants
	once             sync.Once
)

func GetConstants() Constants {
	once.Do(func() {
		port := os.Getenv("PORT")

		if port == "" {
			if err := godotenv.Load(); err != nil {
				panic("Error loading .env file")
			}

			port = ":3000"
		} else {
			port = ":" + port
		}

		constantsInstace = Constants{
			Server: ServerConstants{
				Port: port,
			},
			Database: DatabaseConstants{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				DBName:   os.Getenv("DB_NAME"),
			},
		}
	})

	fmt.Printf("Constants: %v\n", constantsInstace)

	return constantsInstace
}
