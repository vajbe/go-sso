package internal

import (
	"go-sso/internal/types"
	"os"
)

var Conf types.Config

func init() {
	Conf = types.Config{
		Port:        getEnv("USER_SERVICE_PORT", "8080"),
		Db_User:     getEnv("USER_SERVICE_DB_USER", "admin"),
		Db_Pwd:      getEnv("USER_SERVICE_DB_PWD", "admin"),
		Db_Port:     getEnv("USER_SERVICE_DB_PORT", "5432"),
		Db_URL:      getEnv("USER_SERVICE_DB_HOST", "localhost"),
		Redis_URL:   getEnv("USER_SERVICE_REDIS_URL", "localhost:6379"),
		Schema_Path: getEnv("USER_SERVICE_SCHEMA_PATH", "C:\\Users\\vivek\\Documents\\go-microservices\\services\\users\\db\\schema.sql"),
		Kafka_URL:   getEnv("USER_SERVICE_KAFKA_URL", "localhost:9092"),
	}
}

func GetConfig() types.Config {
	return Conf
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
