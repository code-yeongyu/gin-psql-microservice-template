package test

import (
	"fmt"
	"gin_psql_microservice_template/cmd/server/models"
	"gin_psql_microservice_template/cmd/server/utils"
	"gin_psql_microservice_template/configs"
	"testing"

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

func TestDBSetup(t *testing.T) {
	configs.Envs["PSQL_DBNAME"] = "test_database"

	db := utils.GetDB()
	db.SetupDB(models.GetModels())

	resetDB(db.Instance)
}

func TestDBSingleton(t *testing.T) {
	configs.Envs["PSQL_DBNAME"] = "test_database"

	db1 := utils.GetDB()
	db2 := utils.GetDB()
	utils.GetDB().SetupDB(models.GetModels())

	assert.Equal(t, true, db1.Instance == db2.Instance)
}
