// main.go (or cmd/your-app-name/main.go)
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	// If using an internal package, import it like this:
	// "your-quote-app/internal/quotes"
	// "your-quote-app/internal/display"
)

// In a real app, these would be in their own internal packages.
// For demonstration, they are in main.go.

// loadQuotes reads quotes from a file, one per line.
func loadQuotes(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open quotes file: %w", err)
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quote := scanner.Text()
		if len(quote) > 0 { // Avoid empty lines
			quotes = append(quotes, quote)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading quotes file: %w", err)
	}

	return quotes, nil
}

// getRandomQuote selects a random quote from the slice.
func getRandomQuote(quotes []string) (string, error) {
	if len(quotes) == 0 {
		return "", fmt.Errorf("no quotes available")
	}
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex], nil
}

// displayQuote prints the quote to the console.
func displayQuote(quote string) {
	fmt.Println("--- Your Daily Dose of Wisdom ---")
	fmt.Printf("“%s”\n", quote)
	fmt.Println("---------------------------------")
}

func main() {
	// 1. Define Command-Line Flags
	quotesFilePath := flag.String("file", "data/quotes.txt", "Path to the quotes file")
	versionFlag := flag.Bool("version", false, "Print application version")
	flag.Parse() // Parse the flags

	const appVersion = "1.0.0"

	// 2. Handle Flags
	if *versionFlag {
		fmt.Printf("Quote CLI Version: %s\n", appVersion)
		return // Exit after printing version
	}

	// 3. Load Quotes
	quotes, err := loadQuotes(*quotesFilePath)
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}

	// 4. Get a Random Quote
	quote, err := getRandomQuote(quotes)
	if err != nil {
		log.Fatalf("Error getting random quote: %v", err)
	}

	// 5. Display the Quote
	displayQuote(quote)
}
