package search

import (
	"context"
	"time"
)

// DeleteUserByID removes a record from the search table by user ID
func (s *Search) DeleteUserByID(
	UserID uint64,
) error {
	// Construct the SQL query for deletion
	query := "DELETE FROM search WHERE user = ?"

	// Use the passed context to ensure consistent timeout management
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Execute the SQL deletion
	_, err := s.sql.ExecContext(ctx, query, UserID)
	if err != nil {
		return err
	}

	return nil
}
