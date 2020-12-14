package store

import (
	"judaro13/miaguila/storeresultsservice/models"
	"judaro13/miaguila/storeresultsservice/utils"

	"gorm.io/gorm"
)

// SaveUKAPIResponse save data from UKAPI
func SaveUKAPIResponse(db *gorm.DB, data models.UKAPIPOSTResult, reference string) error {
	coordinates := []GeoCoordinate{}
	for _, results := range data.Result {
		for _, value := range results.Result {
			coordinates = append(coordinates, GeoCoordinate{Postcode: value.Postcode,
				Lat: value.Latitude, Lon: value.Longitude})
		}
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	result := tx.CreateInBatches(coordinates, 100)
	if result.Error != nil {
		tx.Rollback()
		utils.Error(result.Error)
		return result.Error
	}

	err := updateStatus(tx, reference)
	if err != nil {
		tx.Rollback()
		utils.Error(result.Error)
		return result.Error
	}

	return tx.Commit().Error
}

func updateStatus(db *gorm.DB, reference string) error {
	progress := CSVUpload{}
	db.Where("reference = ?", reference).First(&progress)
	progress.Counts++
	if progress.Counts >= progress.Bulks {
		progress.Status = "done"
	}

	result := db.Save(&progress)
	if result.Error != nil {
		utils.Error(result.Error)
		return result.Error
	}
	return nil
}
