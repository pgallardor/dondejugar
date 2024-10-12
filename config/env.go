package config

import (
	"fmt"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitConfig() DbConfig {
	return DbConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASS", ""),
		Name:     getEnv("DB_NAME", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func (conf *DbConfig) GetSSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable options='--client_encoding=UTF8'",
		conf.Host, conf.User, conf.Password, conf.Name, conf.Port,
	)
}
