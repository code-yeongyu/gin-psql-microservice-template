package utils

import (
	"fmt"
	"gin_psql_microservice_template/cmd/server/models"
	"gin_psql_microservice_template/configs"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var dbManager *DBManager

// DBManager is a struct to store *gorm.DB
type DBManager struct {
	Instance *gorm.DB
}

// GetDB is a singleton constructor of the DBManager
func GetDB() *DBManager {
	once.Do(func() {
		dbManager = &DBManager{}
	})
	return dbManager
}

// SetupDB setups the gorm DB configuration using the values from the environment
func (m *DBManager) SetupDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.Envs["PSQL_HOST"], configs.Envs["PSQL_USER"], configs.Envs["PSQL_PASSWORD"], configs.Envs["PSQL_DBNAME"], configs.Envs["PSQL_PORT"], configs.Envs["PSQL_SSLMODE"], configs.Envs["PSQL_TIMEZONE"])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	for _, model := range models.GetModels() {
		db.AutoMigrate(model)
	}
	m.Instance = db
}

// InitDB is a shortcut for the GetDB().SetupDB()
func InitDB() {
	GetDB().SetupDB()
}
