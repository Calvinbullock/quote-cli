// internal/quotes/quotes_test.go
package quotes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

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
	if err.Error() != `no quotes found in "` + testFilePath + `"` {
		t.Errorf("Unexpected error message for empty quotes. Got: %q, Want: %q", err.Error(), `no quotes found in "` + testFilePath + `"`)
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
