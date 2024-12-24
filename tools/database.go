package tools

import (
	"database/sql"
	_ "github.com/lib/pq" // Driver para PostgreSQL
)

type Database struct {
	db *sql.DB
}

func NewDatabase(connString string) (*Database, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}
