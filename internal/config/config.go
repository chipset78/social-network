package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	AppPort      string
	JWTSecret    string
	PasswordSalt string
}

func Load() *Config {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	cfg := &Config{}

	// Инициализация флагов с значениями по умолчанию из .env
	flag.StringVar(&cfg.Host, "db-host", getEnv("POSTGRES_HOST", "db"), "Database host")
	flag.IntVar(&cfg.Port, "db-port", getEnvAsInt("POSTGRES_PORT", 5432), "Database port")
	flag.StringVar(&cfg.User, "db-user", getEnv("POSTGRES_USER", "postgres"), "Database user")
	flag.StringVar(&cfg.Password, "db-pass", getEnv("POSTGRES_PASSWORD", "postgres"), "Database password")
	flag.StringVar(&cfg.DBName, "db-name", getEnv("POSTGRES_DB", "social_network"), "Database name")
	flag.StringVar(&cfg.AppPort, "app-port", getEnv("APP_PORT", "8080"), "Application port")
	flag.StringVar(&cfg.JWTSecret, "jwt-secret", getEnv("JWT_SECRET", "default-secret-key"), "JWT secret key")
	flag.StringVar(&cfg.PasswordSalt, "password-salt", getEnv("PASSWORD_SALT", "default-salt"), "Password salt")

	flag.Parse()

	return cfg
}

// Вспомогательные функции для чтения .env
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
