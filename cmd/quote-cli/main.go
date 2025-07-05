package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"quote-cli/internal/quotes"
	//"quote-cli/internal/display"
)

// displayQuote prints the quote to the console.
func displayQuote(quote quotes.Quote) {
	fmt.Println("--- Your Daily Dose of Wisdom ---")
	fmt.Printf("“%s”\n", quote.Text)
	fmt.Printf("  - %s\n", quote.Author)
	fmt.Println("---------------------------------")
}

func main() {
	// Define Command-Line Flags
	quotesFilePath := flag.String("file", "_assets/default.json", "Path to the quotes file")
	versionFlag := flag.Bool("version", false, "Print application version")
	flag.Parse() // Parse the flags

	const appVersion = "1.0.0"

	// Handle Flags
	if *versionFlag {
		fmt.Printf("Quote CLI Version: %s\n", appVersion)
		return // Exit after printing version
	}

	// Load Quotes
	quotes, err := quotes.LoadQuotesFromFile(*quotesFilePath)
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}

	// Display Rand Quote
	randomInt := rand.Intn(len(quotes))
	displayQuote(quotes[randomInt])
}
