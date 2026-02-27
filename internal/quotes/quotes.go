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
// an empty (non-nil) slice of quotes is returned
func SearchByQuoteTag(quotes []Quote, targetTag string, isExact bool) []Quote {
	var matchingQuotes []Quote
	targetTag = strings.ToLower(strings.TrimSpace(targetTag))

	// return quick if empty tag
	if targetTag == "" {
		return matchingQuotes
	}

	// compare tag and targetTag
	for _, quote := range quotes {
		for _, quoteTag := range quote.Tags {
			loweredTag := strings.ToLower(quoteTag)

			if isExact {
				if loweredTag == targetTag {
					matchingQuotes = append(matchingQuotes, quote)
					break
				}
			} else {
				if strings.Contains(loweredTag, targetTag) {
					matchingQuotes = append(matchingQuotes, quote)
					break
				}
			}
		}
	}

	return matchingQuotes
}

// SearchByQuoteAuthor filters a slice of quotes, returning only those written by
// the specified author. The search is case-insensitive and ignores leading/trailing
// whitespace on the authorName.
//
// If the processed authorName is empty, or if no matching quotes are found,
// an empty (non-nil) slice of quotes is returned
func SearchByQuoteAuthor(quotes []Quote, authorName string, isExact bool) []Quote {
	var matchingQuotes []Quote
	authorName = strings.ToLower(strings.TrimSpace(authorName))

	// return quick if empty author
	if authorName == "" {
		return matchingQuotes
	}

	// compare author and targetAuthor
	for _, quote := range quotes {
		loweredAuthor := strings.ToLower(quote.Author)

		// exact or partial author match
		if isExact {
			if loweredAuthor == authorName {
				matchingQuotes = append(matchingQuotes, quote)
			}
		} else {
			if strings.Contains(loweredAuthor, authorName) {
				matchingQuotes = append(matchingQuotes, quote)
			}
		}
	}

	return matchingQuotes
}

// SearchByQuoteAuthor filters a slice of quotes, returning only those written by
// the specified author. The search is case-insensitive and ignores leading/trailing
// whitespace on the authorName.
//
// If the processed authorName is empty, or if no matching quotes are found,
// an empty (non-nil) slice of quotes is returned
func SearchByPartialQuoteAuthor(quotes []Quote, authorName string) []Quote {
	var matchingQuotes []Quote
	authorName = strings.ToLower(strings.TrimSpace(authorName))

	// return quick if empty author
	if authorName == "" {
		return matchingQuotes
	}

	// Pre-allocate memory to avoid multiple small allocations
	matchingQuotes = make([]Quote, 0, len(quotes)/10) // Guessing 10% match rate

	// compare author and targetAuthor
	for _, quote := range quotes {
		if strings.Contains(strings.ToLower(quote.Author), authorName) {
			matchingQuotes = append(matchingQuotes, quote)
		}
	}

	return matchingQuotes
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

// Write Json array to file
func WriteQuoteToFile(quoteList []Quote, filePath string) error {
	// concert to byte slice
	jsonData, err := json.MarshalIndent(quoteList, "", "\t")
	if err != nil {
		return fmt.Errorf("Error marshalling data to JSON: %v\n", err)
	}

	// 4. Write the JSON byte slice to the file
	// os.FileMode(0644) sets the file permissions (read/write for owner, read-only for others).
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing JSON to file %s: %v\n", filePath, err)
	}

	return nil
}

func AddNewQuote(newQuoteText string, author string, tags []string, filePath string) error {
	newQ := Quote{
		Text:   newQuoteText,
		Author: author,
		Tags:   tags,
	}
	quoteList, err := LoadQuotesFromFile(filePath)
	if err != nil {
		return err
	}

	quoteList = append(quoteList, newQ)
	WriteQuoteToFile(quoteList, filePath)

	return nil
}
