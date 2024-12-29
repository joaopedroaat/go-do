package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(homeDir, ".godo", "data", "tasks.db")
	err = os.MkdirAll(filepath.Dir(dbPath), os.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	err = initializeDatabase(db)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func initializeDatabase(db *sql.DB) error {
	stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS Tasks (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      description TEXT NOT NULL,
      done BOOLEAN NOT NULL DEFAULT FALSE
    )
  `)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
