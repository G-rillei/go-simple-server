package constants

import (
	"fmt"
	"os"
	"sync"
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

		fmt.Printf("Port: %v\n", port)

		if port == "" {
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
