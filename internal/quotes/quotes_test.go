// internal/quotes/quotes_test.go
package quotes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

//				Test - SearchByQuoteTag
// ====================================================== \\

// TestSearchByQuoteTag tests the SearchByQuoteTag function.
func TestSearchByQuoteTag(t *testing.T) {
	// Define a set of sample quotes to use across test cases.
	sampleQuotes := []Quote{
		{
			Text:   "The only way to do great work is to love what you do.",
			Author: "Steve Jobs",
			Tags:   []string{"inspiration", "work", "passion"},
		},
		{
			Text:   "Be yourself; everyone else is already taken.",
			Author: "Oscar Wilde",
			Tags:   []string{"identity", "humor"},
		},
		{
			Text:   "The future belongs to those who believe in the beauty of their dreams.",
			Author: "Eleanor Roosevelt",
			Tags:   []string{"future", "dreams", "inspiration"},
		},
		{
			Text:   "Innovation distinguishes between a leader and a follower.",
			Author: "Steve Jobs",
			Tags:   []string{"innovation", "leadership", "work"},
		},
		{
			Text:   "To be or not to be, that is the question.",
			Author: "William Shakespeare",
			Tags:   []string{"philosophy", "drama"},
		},
	}

	// Define test cases. Each test case has a name, input, and expected output.
	tests := []struct {
		name           string
		quotes         []Quote
		targetTag      string
		expectedQuotes []Quote
		expectedError  error // Placeholder for error handling, though current func doesn't return errors
	}{
		{
			name:      "Find single matching quote",
			quotes:    sampleQuotes,
			targetTag: "humor",
			expectedQuotes: []Quote{
				{
					Text:   "Be yourself; everyone else is already taken.",
					Author: "Oscar Wilde",
					Tags:   []string{"identity", "humor"},
				},
			},
			expectedError: nil,
		},
		{
			name:      "Find multiple matching quotes",
			quotes:    sampleQuotes,
			targetTag: "inspiration",
			expectedQuotes: []Quote{
				{
					Text:   "The only way to do great work is to love what you do.",
					Author: "Steve Jobs",
					Tags:   []string{"inspiration", "work", "passion"},
				},
				{
					Text:   "The future belongs to those who believe in the beauty of their dreams.",
					Author: "Eleanor Roosevelt",
					Tags:   []string{"future", "dreams", "inspiration"},
				},
			},
			expectedError: nil,
		},
		{
			name:           "No matching quotes",
			quotes:         sampleQuotes,
			targetTag:      "nonexistent",
			expectedQuotes: []Quote{}, // Expect an empty slice
			expectedError:  nil,
		},
		{
			name:      "Case in-sensitive search",
			quotes:    sampleQuotes,
			targetTag: "Inspiration",
			expectedQuotes: []Quote{
				{
					Text:   "The only way to do great work is to love what you do.",
					Author: "Steve Jobs",
					Tags:   []string{"inspiration", "work", "passion"},
				},
				{
					Text:   "The future belongs to those who believe in the beauty of their dreams.",
					Author: "Eleanor Roosevelt",
					Tags:   []string{"future", "dreams", "inspiration"},
				},
			},
			expectedError: nil,
		},
		{
			name:           "Empty input quotes slice",
			quotes:         []Quote{}, // Empty slice of quotes
			targetTag:      "work",
			expectedQuotes: []Quote{},
			expectedError:  nil,
		},
		{
			name:           "Empty target tag", // Searching for an empty tag
			quotes:         sampleQuotes,
			targetTag:      "",
			expectedQuotes: []Quote{}, // Assuming an empty tag won't match anything
			expectedError:  nil,
		},
		{
			name: "Quote with multiple tags, one matches",
			quotes: []Quote{
				{
					Text:   "Test quote",
					Author: "Tester",
					Tags:   []string{"tag1", "tag2", "tag3"},
				},
			},
			targetTag: "tag2",
			expectedQuotes: []Quote{
				{
					Text:   "Test quote",
					Author: "Tester",
					Tags:   []string{"tag1", "tag2", "tag3"},
				},
			},
			expectedError: nil,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		// Run each test case as a subtest. This helps in organizing test output.
		t.Run(tt.name, func(t *testing.T) {
			// Call the function being tested.
			actualQuotes, actualError := SearchByQuoteTag(tt.quotes, tt.targetTag)

			// Check for errors.
			if actualError != tt.expectedError {
				t.Errorf("SearchByQuoteTag() \nerror   = %v, \nwantErr = %v", actualError, tt.expectedError)
				return // Stop if error expectation is not met
			}

			// Compare the actual returned quotes with the expected quotes.
			// reflect.DeepEqual is used for comparing slices of structs.
			if !reflect.DeepEqual(actualQuotes, tt.expectedQuotes) {
				t.Errorf("SearchByQuoteTag() \ngot  = %v, \nwant = %v", actualQuotes, tt.expectedQuotes)
			}
		})
	}
}

// TestSearchByQuoteAuthor tests the SearchByQuoteAuthor function.
func TestSearchByQuoteAuthor(t *testing.T) {
	// Define a set of sample quotes to use across test cases.
	// This can be the same set used for TestSearchByQuoteTag.
	sampleQuotes := []Quote{
		{
			Text:   "The only way to do great work is to love what you do.",
			Author: "Steve Jobs",
			Tags:   []string{"inspiration", "work", "passion"},
		},
		{
			Text:   "Be yourself; everyone else is already taken.",
			Author: "Oscar Wilde",
			Tags:   []string{"identity", "humor"},
		},
		{
			Text:   "The future belongs to those who believe in the beauty of their dreams.",
			Author: "Eleanor Roosevelt",
			Tags:   []string{"future", "dreams", "inspiration"},
		},
		{
			Text:   "Innovation distinguishes between a leader and a follower.",
			Author: "Steve Jobs",
			Tags:   []string{"innovation", "leadership", "work"},
		},
		{
			Text:   "To be or not to be, that is the question.",
			Author: "William Shakespeare",
			Tags:   []string{"philosophy", "drama"},
		},
		{
			Text:   "All the world's a stage, and all the men and women merely players.",
			Author: "William Shakespeare",
			Tags:   []string{"philosophy", "life"},
		},
	}

	// Define test cases. Each test case has a name, input, and expected output.
	tests := []struct {
		name           string
		quotes         []Quote
		authorName     string
		expectedQuotes []Quote
		expectedError  error
	}{
		{
			name:       "Find single matching quote by author",
			quotes:     sampleQuotes,
			authorName: "Oscar Wilde",
			expectedQuotes: []Quote{
				{
					Text:   "Be yourself; everyone else is already taken.",
					Author: "Oscar Wilde",
					Tags:   []string{"identity", "humor"},
				},
			},
			expectedError: nil,
		},
		{
			name:       "Find multiple matching quotes by author",
			quotes:     sampleQuotes,
			authorName: "Steve Jobs",
			expectedQuotes: []Quote{
				{
					Text:   "The only way to do great work is to love what you do.",
					Author: "Steve Jobs",
					Tags:   []string{"inspiration", "work", "passion"},
				},
				{
					Text:   "Innovation distinguishes between a leader and a follower.",
					Author: "Steve Jobs",
					Tags:   []string{"innovation", "leadership", "work"},
				},
			},
			expectedError: nil,
		},
		{
			name:       "Find multiple matching quotes by author (another example)",
			quotes:     sampleQuotes,
			authorName: "William Shakespeare",
			expectedQuotes: []Quote{
				{
					Text:   "To be or not to be, that is the question.",
					Author: "William Shakespeare",
					Tags:   []string{"philosophy", "drama"},
				},
				{
					Text:   "All the world's a stage, and all the men and women merely players.",
					Author: "William Shakespeare",
					Tags:   []string{"philosophy", "life"},
				},
			},
			expectedError: nil,
		},
		{
			name:           "No matching quotes by author",
			quotes:         sampleQuotes,
			authorName:     "NonExistent Author",
			expectedQuotes: []Quote{}, // Expect an empty slice
			expectedError:  nil,
		},
		{
			name:       "Case insensitive search for author",
			quotes:     sampleQuotes,
			authorName: "steve jobs", // Lowercase input
			expectedQuotes: []Quote{
				{
					Text:   "The only way to do great work is to love what you do.",
					Author: "Steve Jobs",
					Tags:   []string{"inspiration", "work", "passion"},
				},
				{
					Text:   "Innovation distinguishes between a leader and a follower.",
					Author: "Steve Jobs",
					Tags:   []string{"innovation", "leadership", "work"},
				},
			},
			expectedError: nil,
		},
		{
			name:       "Case insensitive search for author (mixed case input)",
			quotes:     sampleQuotes,
			authorName: "OsCaR wIlDe", // Mixed case input
			expectedQuotes: []Quote{
				{
					Text:   "Be yourself; everyone else is already taken.",
					Author: "Oscar Wilde",
					Tags:   []string{"identity", "humor"},
				},
			},
			expectedError: nil,
		},
		{
			name:       "Author name with leading/trailing spaces",
			quotes:     sampleQuotes,
			authorName: "  Steve Jobs  ", // Spaces around the name
			expectedQuotes: []Quote{
				{
					Text:   "The only way to do great work is to love what you do.",
					Author: "Steve Jobs",
					Tags:   []string{"inspiration", "work", "passion"},
				},
				{
					Text:   "Innovation distinguishes between a leader and a follower.",
					Author: "Steve Jobs",
					Tags:   []string{"innovation", "leadership", "work"},
				},
			},
			expectedError: nil,
		},
		{
			name:           "Empty input quotes slice",
			quotes:         []Quote{}, // Empty slice of quotes
			authorName:     "Any Author",
			expectedQuotes: []Quote{},
			expectedError:  nil,
		},
		{
			name:           "Empty author name", // Searching for an empty author name
			quotes:         sampleQuotes,
			authorName:     "",
			expectedQuotes: []Quote{}, // Expect an empty slice
			expectedError:  nil,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		// Run each test case as a subtest.
		t.Run(tt.name, func(t *testing.T) {
			// Call the function being tested.
			actualQuotes, actualError := SearchByQuoteAuthor(tt.quotes, tt.authorName)

			// Check for errors.
			if actualError != tt.expectedError {
				t.Errorf("SearchByQuoteAuthor() error = %v, wantErr %v", actualError, tt.expectedError)
				return
			}

			// Compare the actual returned quotes with the expected quotes.
			// reflect.DeepEqual is used for comparing slices of structs.
			if !reflect.DeepEqual(actualQuotes, tt.expectedQuotes) {
				t.Errorf("SearchByQuoteAuthor() got = %v, want %v", actualQuotes, tt.expectedQuotes)
			}
		})
	}
}

//				Test - LoadQuotesFromFile
// ====================================================== \\

// TestLoadQuotesFromFile_Success ... (unchanged)
func TestLoadQuotesFromFile_Success(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "valid_quotes.json")

	validJSON := `[
	{"text": "Test Quote 1", "author": "Test Author 1"},
	{"text": "Test Quote 2", "author": "Test Author 2"}
	]`
	err := os.WriteFile(testFilePath, []byte(validJSON), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	expectedQuotes := []Quote{
		{Text: "Test Quote 1", Author: "Test Author 1"},
		{Text: "Test Quote 2", Author: "Test Author 2"},
	}

	quotes, err := LoadQuotesFromFile(testFilePath)

	if err != nil {
		t.Fatalf("LoadQuotesFromFile returned an unexpected error: %v", err)
	}
	if !reflect.DeepEqual(quotes, expectedQuotes) {
		t.Errorf("LoadQuotesFromFile returned incorrect quotes.\nGot: %+v\nWant: %+v", quotes, expectedQuotes)
	}
}

// TestLoadQuotesFromFile_FileNotFound ... (unchanged)
func TestLoadQuotesFromFile_FileNotFound(t *testing.T) {
	nonExistentPath := "non_existent_file.json"

	quotes, err := LoadQuotesFromFile(nonExistentPath)

	if err == nil {
		t.Error("LoadQuotesFromFile expected an error for non-existent file, but got none.")
	}
	if quotes != nil {
		t.Errorf("LoadQuotesFromFile expected nil quotes for non-existent file, but got: %+v", quotes)
	}
}

// TestLoadQuotesFromFile_EmptyQuotesArray ... (unchanged)
func TestLoadQuotesFromFile_EmptyQuotesArray(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "empty_quotes.json")
	err := os.WriteFile(testFilePath, []byte("[]"), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	quotes, err := LoadQuotesFromFile(testFilePath)

	if err == nil {
		t.Error("LoadQuotesFromFile expected an error for empty quotes array, but got none.")
	}
	if quotes != nil {
		t.Errorf("LoadQuotesFromFile expected nil quotes for empty array, but got: %+v", quotes)
	}
	if err.Error() != `no quotes found in "`+testFilePath+`"` {
		t.Errorf("Unexpected error message for empty quotes. Got: %q, Want: %q", err.Error(), `no quotes found in "`+testFilePath+`"`)
	}
}

// TestLoadQuotesFromFile_MalformedJSON tests invalid JSON without 'errors.As'.
func TestLoadQuotesFromFile_MalformedJSON(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "malformed_quotes.json")
	malformedJSON := `[{"text": "Bad quote", "author": "Bad Author"}, {"text": "Missing comma" "author": "Another Bad Author"}]`

	err := os.WriteFile(testFilePath, []byte(malformedJSON), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	quotes, err := LoadQuotesFromFile(testFilePath)

	if err == nil {
		t.Error("LoadQuotesFromFile expected an error for malformed JSON, but got none.")
	}
	if quotes != nil {
		t.Errorf("LoadQuotesFromFile expected nil quotes for malformed JSON, but got: %+v", quotes)
	}

	// --- Manual Error Unwrapping and Type Assertion ---
	// Start with the top-level error returned by LoadQuotesFromFile
	currentErr := err
	foundSyntaxError := false
	for currentErr != nil {
		if _, ok := currentErr.(*json.SyntaxError); ok {
			foundSyntaxError = true
			break // Found the specific error type
		}
		// If the current error can be unwrapped, get the next error in the chain
		// Use a type assertion to the error interface and then check for Unwrap method.
		// fmt.Errorf with %w creates errors that implement `Unwrap() error`.
		if unwrapper, ok := currentErr.(interface{ Unwrap() error }); ok {
			currentErr = unwrapper.Unwrap()
		} else {
			// No more errors to unwrap
			break
		}
	}

	if !foundSyntaxError {
		t.Errorf("Expected an error containing *json.SyntaxError, but did not find it in error chain: %v", err)
	}
}

// TestLoadQuotesFromFile_InvalidJSONType tests when JSON is not an array without 'errors.As'.
func TestLoadQuotesFromFile_InvalidJSONType(t *testing.T) {
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "invalid_type.json")
	invalidTypeJSON := `{"text": "Not an array"}`

	err := os.WriteFile(testFilePath, []byte(invalidTypeJSON), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	quotes, err := LoadQuotesFromFile(testFilePath)

	if err == nil {
		t.Error("LoadQuotesFromFile expected an error for invalid JSON type, but got none.")
	}
	if quotes != nil {
		t.Errorf("LoadQuotesFromFile expected nil quotes for invalid JSON type, but got: %+v", quotes)
	}

	// --- Manual Error Unwrapping and Type Assertion ---
	currentErr := err
	foundUnmarshalTypeError := false
	for currentErr != nil {
		if _, ok := currentErr.(*json.UnmarshalTypeError); ok {
			foundUnmarshalTypeError = true
			break // Found the specific error type
		}
		if unwrapper, ok := currentErr.(interface{ Unwrap() error }); ok {
			currentErr = unwrapper.Unwrap()
		} else {
			break
		}
	}

	if !foundUnmarshalTypeError {
		t.Errorf("Expected an error containing *json.UnmarshalTypeError, but did not find it in error chain: %v", err)
	}
}
