// Package db contains helpers to open an SQLite database and perform schema migration.
package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// OpenAndMigrate opens an SQLite database at the given path and ensures the schema exists.
// It will create parent directories as needed.
func OpenAndMigrate(dbPath string) (*sql.DB, error) {
	// Ensure parent directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("create db dir: %w", err)
	}

	dsn := fmt.Sprintf("file:%s?_pragma=busy_timeout(5000)", dbPath)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	// Create schema if not exists
	if err := migrate(db); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

// migrate ensures required tables exist.
func migrate(db *sql.DB) error {
	const createItems = `
CREATE TABLE IF NOT EXISTS items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ','now'))
);
`
	if _, err := db.Exec(createItems); err != nil {
		return fmt.Errorf("create items: %w", err)
	}
	return nil
}
