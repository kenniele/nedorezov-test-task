package config

import "os"

type DBConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMODE  string
}

type Config struct {
	DB DBConfig
}

func New() *Config {
	return &Config{
		DB: DBConfig{
			HOST:     getEnv("DB_HOST", "localhost"),
			PORT:     getEnv("DB_PORT", "5432"),
			USER:     getEnv("DB_USER", "postgres"),
			PASSWORD: getEnv("DB_PASSWORD", "postgres"),
			DBNAME:   getEnv("DB_NAME", "postgres"),
			SSLMODE:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
