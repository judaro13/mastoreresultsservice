package process

import (
	"encoding/json"
	"judaro13/miaguila/storeresultsservice/store"

	"github.com/judaro13/masharedmodels/models"

	"gorm.io/gorm"
)

// Data process data for store
func Data(db *gorm.DB, data []byte) error {
	message := models.StoreDataMessage{}
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}

	store.SaveUKAPIResponse(db, message.Result, message.Reference)
	if err != nil {
		return err
	}

	return nil
}
