package store

import (
	"judaro13/miaguila/storeresultsservice/testshelper"
	"testing"

	"github.com/judaro13/masharedmodels/models"

	"github.com/stretchr/testify/assert"
)

func TestUpdateStatus(t *testing.T) {
	db, err := testshelper.PrepareTestDB()
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
	db, err := testshelper.PrepareTestDB()
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
