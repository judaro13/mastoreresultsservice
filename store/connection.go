package store

import (
	"fmt"
	"os"

	"github.com/judaro13/masharedmodels/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDB func that create a DB connection
func ConnectToDB() *gorm.DB {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Connect DB", r)
		}
	}()

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Automigrations(db)
	return db
}

func Automigrations(db *gorm.DB) {
	db.AutoMigrate(&models.GeoCoordinate{})
	db.AutoMigrate(&models.CSVUpload{})
}
