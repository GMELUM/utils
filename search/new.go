package search

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Search struct {
	sql       *sql.DB
	interests map[string]bool
}

func New(config ...Config) (*Search, error) {

	cfg := configDefault(config...)

	// Remove existing database file, if any

	if cfg.Reset && cfg.LocalFile != "file::memory:" {
		err := os.Remove(cfg.LocalFile)
		if err != nil && !os.IsNotExist(err) {
			return nil, err // Return an error if it's not "file does not exist" error
		}
	}

	// Initialize SQLite database connection
	db, err := sql.Open("sqlite3", cfg.LocalFile)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is valid.
	err = db.Ping()
	if err != nil {
		return nil, err // Terminate the program if the connection is invalid.
	}

	if err := create(db, cfg.Interests); err != nil {
		return nil, err
	}

	interestsMap := make(map[string]bool, len(cfg.Interests))
	for _, interest := range cfg.Interests {
		interestsMap[interest] = true
	}

	search := &Search{
		sql:       db,
		interests: interestsMap,
	}

	return search, nil

}

type Params struct {
	Query string // SQL query string
	Args  []any  // Arguments for the SQL query
}

func query[T any](
	s *Search,
	params Params,
	callback func(rows *sql.Rows) (*T, error),
) (*T, error) {

	// Create a context with a timeout for the query execution
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel() // Cancel the context after the query execution

	// Execute the query with the provided arguments
	rows, err := s.sql.QueryContext(ctx, params.Query, params.Args...)
	if err != nil {
		// Return the SQL error if it is any other error
		return nil, err
	}
	defer rows.Close() // Close the rows after finishing the query

	// Call the callback function to process the rows and extract the result
	clbRes, clbErr := callback(rows)

	// Return the result and any potential MySQL error from the callback
	return clbRes, clbErr

}

func (s *Search) Close() error {
	return s.sql.Close()
}
