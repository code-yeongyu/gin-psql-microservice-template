package utils

import (
	"fmt"
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

func getDsn() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.Envs["PSQL_HOST"], configs.Envs["PSQL_USER"], configs.Envs["PSQL_PASSWORD"], configs.Envs["PSQL_DBNAME"], configs.Envs["PSQL_PORT"], configs.Envs["PSQL_SSLMODE"], configs.Envs["PSQL_TIMEZONE"])
	return dsn
}

// GetDB is a singleton constructor of the DBManager
func GetDB() *DBManager {
	once.Do(func() {
		dbManager = &DBManager{}
	})
	return dbManager
}

// OpenDB opens db
func OpenDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// MigrateModels automatically migrates registered models
func MigrateModels(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		db.AutoMigrate(model)
	}
}

// SetupDB setups the gorm DB configuration using the values from the environment
func (m *DBManager) SetupDB(models []interface{}) {
	dsn := getDsn()
	db := OpenDB(dsn)
	MigrateModels(db, models)
	m.Instance = db
}

// InitDB is a shortcut for the GetDB().SetupDB()
func InitDB(models []interface{}) {
	GetDB().SetupDB(models)
}
