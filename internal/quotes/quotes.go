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

// read json data from file, return as slice of Quotes
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
