package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/data/tasks.db")
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
