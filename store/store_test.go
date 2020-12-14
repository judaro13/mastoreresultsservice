package store

import (
	"os"
	"testing"

	"github.com/judaro13/masharedmodels/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PrepareTestDB prepare test DB according to txdb name
func PrepareTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	Automigrations(db)
	return db, err
}

func TestUpdateStatus(t *testing.T) {
	db, err := PrepareTestDB()
	assert.NoError(t, err)
	tx := db.Begin()

	err = updateStatus(tx, "asdf")
	assert.Error(t, err)

	status := models.CSVUpload{Reference: "asdf"}
	result := tx.Create(&status)
	assert.NoError(t, result.Error)

	err = updateStatus(tx, "asdf")
	assert.NoError(t, err)
	tx.Rollback()
}

func TestSaveUKAPIResponse(t *testing.T) {
	db, err := PrepareTestDB()
	assert.NoError(t, err)
	tx := db.Begin()

	data := models.UKAPIPOSTResult{
		Result: []models.UKAPIResults{
			models.UKAPIResults{
				Result: []models.UKAPICoordinate{
					models.UKAPICoordinate{Postcode: "ASD 75", Latitude: 123, Longitude: 321},
				},
			},
		},
	}

	err = SaveUKAPIResponse(tx, data, "asdf")
	assert.Error(t, err)
	tx.Rollback()
}
