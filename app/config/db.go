package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewEnvDBConfig() *DBConfig {
	return &DBConfig{
		Host:     getEnv("DB_HOST", "db"),
		Port:     getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USERNAME", "default"),
		Password: getEnv("DB_PASSWORD", "example"),
		Database: getEnv("DB_DATABASE", "db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
