package search

import (
	"testing"
)

// Test case for searching a suitable interlocutor without specific interests.
func TestSearch(t *testing.T) {
	search, err := New(Config{
		Interests: []string{
			"music", "travel", "sport", "art", "cooking", "movies", "games",
			"reading", "tech", "animals", "nature", "photography", "dance",
			"space", "science", "history", "fashion", "yoga", "psychology",
			"volunteering", "flirt", "crypto", "anime", "lgbt",
		},
	})
	if err != nil {
		t.Fatalf("Error during initialization: %v", err)
	}
	defer search.Close()

	// Assume that some entries were created previously
	search.Create(1, "en", 18, 30, 1, 25, 1, "music", "travel", "art")
	search.Create(2, "en", 18, 30, 1, 25, 1, "movies", "science", "tech")

	result, err := search.Search("en", 18, 30, 1, 1, 25)
	if err != nil {
		t.Fatalf("Error during search: %v", err)
	}
	if result == nil || result.UserID != 1 {
		t.Errorf("Expected to find user 1, got %+v", result)
	}
}

// Test case for searching a suitable interlocutor with specific interests.
func TestSearchWithInterests(t *testing.T) {
	search, err := New(Config{
		Interests: []string{
			"music", "travel", "sport", "art", "cooking", "movies", "games",
			"reading", "tech", "animals", "nature", "photography", "dance",
			"space", "science", "history", "fashion", "yoga", "psychology",
			"volunteering", "flirt", "crypto", "anime", "lgbt",
		},
	})
	if err != nil {
		t.Fatalf("Error during initialization: %v", err)
	}
	defer search.Close()

	// Assume that some entries were created previously
	search.Create(1, "en", 18, 30, 1, 25, 1, "music", "travel", "art")
	search.Create(2, "en", 18, 30, 1, 25, 1, "movies", "science", "tech")
	search.Create(3, "en", 18, 30, 1, 25, 1, "science", "photography")

	result, err := search.Search("en", 18, 30, 1, 1, 25, "music", "art")
	if err != nil {
		t.Fatalf("Error during search: %v", err)
	}
	if result == nil || result.UserID != 1 {
		t.Errorf("Expected to find user 1 that matches the interests 'music' and 'art', got %+v", result)
	}
}

// Benchmarks for the Search function with specific interests
func BenchmarkSearchWithInterests(b *testing.B) {
	search, err := New(Config{
		Reset: false,
		Interests: []string{
			"music", "travel", "sport", "art", "cooking", "movies", "games",
			"reading", "tech", "animals", "nature", "photography", "dance",
			"space", "science", "history", "fashion", "yoga", "psychology",
			"volunteering", "flirt", "crypto", "anime", "lgbt",
		},
	})
	if err != nil {
		b.Fatalf("Error during initialization: %v", err)
	}
	defer search.Close()

	// Insert records to search for
	err = search.Create(1, "en", 18, 30, 1, 25, 1, "music", "travel", "art")
	if err != nil {
		b.Fatalf("Error during record creation for benchmark: %v", err)
	}

	err = search.Create(2, "en", 18, 30, 1, 25, 1, "movies", "science", "tech")
	if err != nil {
		b.Fatalf("Error during record creation for benchmark: %v", err)
	}

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_, err := search.Search("en", 18, 30, 1, 1, 25, "music", "art")
		if err != nil {
			b.Fatalf("Error during search with interests: %v", err)
		}
	}
}

// Benchmarks for the Search function without specific interests
func BenchmarkSearch(b *testing.B) {
	search, err := New(Config{
		Reset: false,
		Interests: []string{
			"music", "travel", "sport", "art", "cooking", "movies", "games",
			"reading", "tech", "animals", "nature", "photography", "dance",
			"space", "science", "history", "fashion", "yoga", "psychology",
			"volunteering", "flirt", "crypto", "anime", "lgbt",
		},
	})
	if err != nil {
		b.Fatalf("Error during initialization: %v", err)
	}
	defer search.Close()

	// Insert a record to search for
	err = search.Create(1, "en", 18, 30, 1, 25, 1, "music", "travel", "art")
	if err != nil {
		b.Fatalf("Error during record creation for benchmark: %v", err)
	}

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_, err := search.Search("en", 18, 30, 1, 1, 25)
		if err != nil {
			b.Fatalf("Error during search: %v", err)
		}
	}
}
