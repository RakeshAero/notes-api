package database

import ( 
	"database/sql" 
	"fmt"
	_ "modernc.org/sqlite" 
)

func Open(dsn string) (*sql.DB, error){
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("Opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("pinging database : %w", err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_IMESTAMP
	);
	`

	notesTable := ` CREATE TABLE IF NOT EXISTS notes ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		user_id INTEGER NOT NULL, 
		title TEXT NOT NULL, 
		content TEXT NOT NULL, 
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		FOREIGN KEY (user_id) 
		REFERENCES users(id) 
	);`

	if _, err := db.Exec(usersTable); err != nil { 
		return fmt.Errorf("creating users table: %w", err) 
	}

	if _, err := db.Exec(notesTable); err != nil { 
		return fmt.Errorf("creating notes table: %w", err) 
	}

	return nil
}