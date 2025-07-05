package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"quote-cli/internal/quotes"
	//"quote-cli/internal/display"
)

// displayQuoteList prints a list of quotes to the console.
func displayQuoteList(quoteList []quotes.Quote) {
	for _, quote := range quoteList {
		displayQuote(quote)
	}
}

// displayQuote prints the quote to the console.
func displayQuote(quote quotes.Quote) {
	fmt.Println("")
	fmt.Printf("“%s”\n", quote.Text)
	fmt.Printf("  - %s\n", quote.Author)
	fmt.Println("")
}

func main() {
	// Define Command-Line Flags
	quotesFilePath := flag.String("file", "_assets/default.json", "Path to the quotes file")
	quotesTagSearch := flag.String("tag", "", "Tag to search quotes for")
	versionFlag := flag.Bool("version", false, "Print application version")
	flag.Parse() // Parse the flags

	const appVersion = "1.0.0"

	// Display program version
	if *versionFlag {
		fmt.Printf("Quote CLI Version: %s\n", appVersion)
		return
	}

	// Load Quotes
	quoteList, err := quotes.LoadQuotesFromFile(*quotesFilePath)
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}

	if *quotesTagSearch != "" {
		// Handle tag search
		foundQuotes, err := quotes.SearchByQuoteTag(quoteList, *quotesTagSearch)
		if err != nil {
			log.Fatalf("Error with search: %v", err)
		}
		displayQuoteList(foundQuotes)
	} else {
		// Display Random Quote
		randomInt := rand.Intn(len(quoteList))
		displayQuote(quoteList[randomInt])
	}
}
