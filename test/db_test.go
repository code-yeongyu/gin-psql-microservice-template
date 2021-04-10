package test

import (
	"fmt"
	"testing"
	"user_service/cmd/server/models"
	"user_service/cmd/server/utils"
	"user_service/configs"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tj/assert"
	"gorm.io/gorm"
)

func resetDB(db *gorm.DB) {
	for _, model := range models.GetModels() {
		modelName := utils.GetTableName(model)
		db.Exec(fmt.Sprintf("DELETE FROM %s", modelName)) // reset db data
	}
}

func testClearDB(t *testing.T) {
	db := utils.GetDB().Instance
	resetDB(db)
}

func TestDBSingleton(t *testing.T) {
	db1 := utils.GetDB()
	db2 := utils.GetDB()
	utils.GetDB().SetupDB(models.GetModels())

	assert.Equal(t, true, db1.Instance == db2.Instance)
}

func TestDBSetup(t *testing.T) {
	configs.Envs["PSQL_DBNAME"] = "test_database"

	db := utils.GetDB()
	db.SetupDB(models.GetModels())

	resetDB(db.Instance)
}
