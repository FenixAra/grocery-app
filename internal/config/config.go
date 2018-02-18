package config

import (
	"os"
	"strconv"
)

var (
	APP_NAME            string
	MIGRATION_FILE_PATH string
	PORT                string
	DMS_URL             string
	SERVICE_TOKEN       string

	DATABASE_URL       string
	DATABASE_URL_TEST  string
	MAX_DB_CONNECTIONS int

	REDIS_URL                         string
	REDIS_CONN_POOL_IDLE_TIMEOUT_MINS int
	REDIS_MAX_ACTIVE_CONNECTIONS      int
	EXPIRY_TIME                       int
)

func init() {
	APP_NAME = os.Getenv("APP_NAME")
	DMS_URL = os.Getenv("DMS_URL")
	MIGRATION_FILE_PATH = os.Getenv("MIGRATION_FILE_PATH")
	PORT = os.Getenv("PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	DATABASE_URL_TEST = os.Getenv("DATABASE_URL_TEST")
	MAX_DB_CONNECTIONS, _ = strconv.Atoi(os.Getenv("MAX_DB_CONNECTIONS"))
	SERVICE_TOKEN = os.Getenv("SERVICE_TOKEN")

	REDIS_URL = os.Getenv("REDIS_URL")
	REDIS_CONN_POOL_IDLE_TIMEOUT_MINS, _ = strconv.Atoi(os.Getenv("REDIS_CONN_POOL_IDLE_TIMEOUT_MINS"))
	REDIS_MAX_ACTIVE_CONNECTIONS, _ = strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE_CONNECTIONS"))
	EXPIRY_TIME, _ = strconv.Atoi(os.Getenv("EXPIRY_TIME"))
}
