package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig database connection details
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// New creates a new Config with values from environment variables or defaults
func New() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DATABASE_HOST", "db"),
			Port:     getEnv("DATABASE_PORT", "5432"),
			User:     getEnv("DATABASE_USER", "postgres"),
			Password: getEnv("DATABASE_PASSWORD", "postgres"),
			DBName:   getEnv("DATABASE_NAME", "test_db"),
		},
	}
}

// DSN returns a formatted database connection string
func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)
}

// getEnv retrieves an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
