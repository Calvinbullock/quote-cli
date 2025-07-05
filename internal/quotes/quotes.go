package quotes

import (
	"encoding/json"
	"fmt"
	"os"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

// LoadQuotesFromFile reads a JSON file from the given filepath,
// parses its content, and returns a slice of Quote structs.
//
// The function expects the JSON file to contain an array of objects,
// where each object can be unmarshaled into a Quote struct.
//
// It returns an error if:
//   - The file cannot be read (e.g., due to non-existence or permissions).
//   - The file content is not valid JSON or cannot be unmarshaled into []Quote.
//   - The JSON file is valid but contains an empty array.
func LoadQuotesFromFile(filepath string) ([]Quote, error) {
	// Read the entire file content
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read quotes file %q: %w", filepath, err)
	}

	var quotes []Quote
	// pares json into struct
	err = json.Unmarshal(data, &quotes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from %q: %w", filepath, err)
	}

	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quotes found in %q", filepath)
	}

	return quotes, nil
}
