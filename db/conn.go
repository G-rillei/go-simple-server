package db

import (
	"fmt"
	"sync"

	"github.com/G-rillei/go-simple-server/constants"
	"github.com/G-rillei/go-simple-server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		constants := constants.GetConstants()

		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			constants.Database.Host,
			constants.Database.Port,
			constants.Database.User,
			constants.Database.DBName,
			constants.Database.Password,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Failed to connect to database!")
		}

		db.AutoMigrate(&model.Post{})

		dbInstance = db
	})

	return dbInstance
}
