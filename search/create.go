package search

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Create inserts a new record into the search table
func Create(
	UserID uint64,
	Language string,
	YourStart int,
	YourEnd int,
	YourSex int,
	MyAge int,
	MySex int,
	interests ...string,
) error {

	// Check if core is initialized
	if core == nil || core.sql == nil {
		return ErrNotInitialize
	}

	// Base columns
	columns := []string{
		"user", "language", "your_start", "your_end", "your_sex", "my_age", "my_sex",
	}

	// Base values
	values := []interface{}{
		UserID, Language, YourStart, YourEnd, YourSex, MyAge, MySex,
	}

	// Process and validate interests
	for _, interest := range interests {
		if core.interests[interest] { // Ensure interest is valid
			columns = append(columns, interest)
			values = append(values, 1) // Assume the presence of interest should be set to 1
		}
	}

	// Create column string and value placeholders
	columnsStr := strings.Join(columns, ", ")
	placeholders := strings.TrimRight(strings.Repeat("?, ", len(values)), ", ")

	// Construct the SQL query
	query := fmt.Sprintf("INSERT INTO search (%s) VALUES (%s)", columnsStr, placeholders)

	// Execute the query with context to manage execution time
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := core.sql.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
