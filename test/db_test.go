package test

import (
	"fmt"
	"gin_psql_microservice_template/cmd/server/models"
	"gin_psql_microservice_template/cmd/server/utils"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tj/assert"
)

func TestDBSingleton(t *testing.T) {
	db1 := utils.GetDB()
	db2 := utils.GetDB()
	utils.GetDB().SetupDB()

	assert.Equal(t, true, db1.Instance == db2.Instance)
}

func TestDBSetup(t *testing.T) {
	db := utils.GetDB()
	db.SetupDB()
	for _, model := range models.GetModels() {
		modelName := utils.GetType(model)
		db.Instance.Exec(fmt.Sprintf("DELETE FROM %s", modelName)) // reset db data
	}
}
