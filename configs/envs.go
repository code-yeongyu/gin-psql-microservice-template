package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func ternaryMap(value string, defaultValue string) string {
	return map[bool]string{true: value, false: defaultValue}[len(value) != 0]
}

func getEnv(envName string, defaultValue string) string {
	return ternaryMap(os.Getenv(envName), defaultValue)
}

// Envs has values for environment variables and the defaults for them
var Envs = map[string]string{
	"ENABLE_SWAGGER": getEnv("IS_ENABLE_SWAGGER", "true"),
	"SERVER_PORT":    getEnv("PORT", "8080"),
	"PSQL_HOST":      getEnv("PSQL_HOST", "localhost"),
	"PSQL_PORT":      getEnv("PSQL_PORT", "5432"),
	"PSQL_USER":      getEnv("PSQL_USER", "postgres"),
	"PSQL_PASSWORD":  getEnv("PSQL_PASSWORD", "1234"),
	"PSQL_DBNAME":    getEnv("PSQL_DBNAME", "database"),
	"PSQL_TIMEZONE":  getEnv("PSQL_TIMEZONE", "Asia/Seoul"),
	"PSQL_SSLMODE":   getEnv("PSQL_SSLMODE", "disable"),
}
