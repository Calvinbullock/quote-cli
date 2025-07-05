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

	if quote.Author == "" {
		fmt.Printf("  - ? \n")
	} else {
		fmt.Printf("  - %s\n", quote.Author)
	}
	fmt.Println("")
}

func main() {
	var quotesFilePath string
	var quotesTagSearch string
	var quotesAuthorSearch string
	var versionFlag bool

	// src file
	flag.StringVar(&quotesFilePath, "file", "_assets/default.json", "Path to the quotes file")
	flag.StringVar(&quotesFilePath, "f", "_assets/default.json", "Path to the quotes file")
	//tag
	flag.StringVar(&quotesTagSearch, "tag", "", "Tag to search quotes for (case-insensitive)")
	flag.StringVar(&quotesTagSearch, "t", "", "Tag to search quotes for (case-insensitive)")
	// author
	flag.StringVar(&quotesAuthorSearch, "author", "", "Author to search quotes for (case-insensitive)")
	flag.StringVar(&quotesAuthorSearch, "a", "", "Short for --author (case-insensitive)")
	// version
	flag.BoolVar(&versionFlag, "version", false, "Print application version")
	flag.BoolVar(&versionFlag, "v", false, "Print application version")
	flag.Parse()

	const appVersion = "1.0.0"

	// Display program version
	if versionFlag {
		fmt.Printf("Quote CLI Version: %s\n", appVersion)
		return
	}

	// Load Quotes
	quoteList, err := quotes.LoadQuotesFromFile(quotesFilePath)
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}

	if quotesTagSearch != "" {
		// Handle tag search
		foundQuotes, err := quotes.SearchByQuoteTag(quoteList, quotesTagSearch)
		if err != nil {
			log.Fatalf("Error with search: %v", err)
		}
		displayQuoteList(foundQuotes)

	} else if quotesAuthorSearch != "" {
		// Handle Author search
		foundQuotes, err := quotes.SearchByQuoteAuthor(quoteList, quotesAuthorSearch)
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
