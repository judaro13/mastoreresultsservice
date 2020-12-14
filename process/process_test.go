package process

import (
	"judaro13/miaguila/storeresultsservice/testshelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateStatus(t *testing.T) {
	db, err := testshelper.PrepareTestDB()
	assert.NoError(t, err)
	tx := db.Begin()

	err = Data(tx, []byte("asdf"))
	assert.Error(t, err)
	tx.Rollback()
}
