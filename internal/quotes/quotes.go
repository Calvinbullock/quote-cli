package quotes

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Quote struct {
	Text   string   `json:"text"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

// SearchByQuoteTag filters a slice of quotes, returning only those that contain
// the specified targetTag. The search is case-insensitive and ignores leading/trailing
// whitespace on the targetTag.
//
// If the processed targetTag is empty, or if no matching quotes are found,
// an empty (non-nil) slice of quotes is returned along with a nil error.
// Errors are reserved for unexpected issues during the search process itself.
func SearchByQuoteTag(quotes []Quote, targetTag string) ([]Quote, error) {
	var matchingQuotes []Quote
	targetTag = strings.ToLower(strings.TrimSpace(targetTag))

	// return quick if empty tag
	if targetTag == "" {
		return matchingQuotes, nil
	}

	// compare tag and targetTag
	for _, quote := range quotes {
		for _, quoteTag := range quote.Tags {
			if strings.ToLower(quoteTag) == targetTag {
				matchingQuotes = append(matchingQuotes, quote)
				break
			}
		}
	}

	return matchingQuotes, nil
}

// SearchByQuoteAuthor filters a slice of quotes, returning only those written by
// the specified author. The search is case-insensitive and ignores leading/trailing
// whitespace on the authorName.
//
// If the processed authorName is empty, or if no matching quotes are found,
// an empty (non-nil) slice of quotes is returned along with a nil error.
// Errors are reserved for unexpected issues during the search process itself.
func SearchByQuoteAuthor(quotes []Quote, authorName string) ([]Quote, error) {
	var matchingQuotes []Quote
	authorName = strings.ToLower(strings.TrimSpace(authorName))

	// return quick if empty author
	if authorName == "" {
		return matchingQuotes, nil
	}

	// compare author and targetAuthor
	for _, quote := range quotes {
		if strings.ToLower(quote.Author) == authorName {
			matchingQuotes = append(matchingQuotes, quote)
		}
	}

	return matchingQuotes, nil
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
	// TODO: change tag slices to maps for quick look ups?

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
