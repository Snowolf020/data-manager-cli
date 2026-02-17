package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgresConfig represents the configuration for a PostgreSQL database connection.
type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// PostgresDB is a PostgreSQL database connection handler.
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB returns a new PostgreSQL database connection handler.
func NewPostgresDB(config PostgresConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DBName)

db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

// Close closes the PostgreSQL database connection.
func (p *PostgresDB) Close() error {
	return p.db.Close()
}

// Query executes a SQL query on the PostgreSQL database.
func (p *PostgresDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.Query(query, args...)
}

// Exec executes a SQL query on the PostgreSQL database.
func (p *PostgresDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return p.db.Exec(query, args...)
}

// Begin starts a new transaction on the PostgreSQL database.
func (p *PostgresDB) Begin() (*sql.Tx, error) {
	return p.db.Begin()
}

func main() {}
