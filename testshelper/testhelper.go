package testshelper

import (
	"os"

	"github.com/judaro13/masharedmodels/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PrepareTestDB prepare test DB according to txdb name
func PrepareTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	db.AutoMigrate(&models.GeoCoordinate{})
	db.AutoMigrate(&models.CSVUpload{})
	return db, err
}
