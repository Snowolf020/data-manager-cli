package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	Database      DatabaseConfig      `json:"database"`
	Prisma        PrismaConfig        `json:"prisma"`
	Redis         RedisConfig         `json:"redis"`
	Docker        DockerConfig        `json:"docker"`
	SQLite        SQLiteConfig        `json:"sqlite"`
	PostgreSQL    PostgreSQLConfig    `json:"postgresql"`
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Type        string `json:"type"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
}

// PrismaConfig represents the Prisma configuration
type PrismaConfig struct {
	Schema      string `json:"schema"`
}

// RedisConfig represents the Redis configuration
type RedisConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
}

// DockerConfig represents the Docker configuration
type DockerConfig struct {
	ComposeFile string `json:"composeFile"`
}

// SQLiteConfig represents the SQLite configuration
type SQLiteConfig struct {
	Database    string `json:"database"`
}

// PostgreSQLConfig represents the PostgreSQL configuration
type PostgreSQLConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
}

// LoadConfig loads the application configuration from a file
func LoadConfig(filePath string) (*Config, error) {
	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// ValidateConfig validates the application configuration
func ValidateConfig(config *Config) error {
	if config.Database.Type == "" {
		return errors.New("database type is required")
	}

	if config.Prisma.Schema == "" {
		return errors.New("Prisma schema is required")
	}

	if config.Redis.Host == "" || config.Redis.Port == 0 {
		return errors.New("Redis host and port are required")
	}

	if config.Docker.ComposeFile == "" {
		return errors.New("Docker compose file is required")
	}

	if config.SQLite.Database == "" {
		return errors.New("SQLite database is required")
	}

	if config.PostgreSQL.Host == "" || config.PostgreSQL.Port == 0 || config.PostgreSQL.Username == "" || config.PostgreSQL.Password == "" || config.PostgreSQL.Name == "" {
		return errors.New("PostgreSQL host, port, username, password, and name are required")
	}

	return nil
}
