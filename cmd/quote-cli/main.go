package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"quote-cli/internal/display"
	"quote-cli/internal/quotes"
)

const appVersion = "1.0.0"

// path for testing
const testFilePath = "_assets/small.json"

// path for real build
const appConfigRelativePath = "quote-cli"
const configFileName = "default.json"

// getDefaultConfigPath returns the full path to the default configuration file
// in an OS-idiomatic location.
func getDefaultConfigPath() (string, error) {
	// os.UserConfigDir() provides the OS-specific user configuration directory:
	// - Linux:   $XDG_CONFIG_HOME or $HOME/.config
	// - macOS:   $HOME/Library/Application Support
	// - Windows: %APPDATA% (e.g., C:\Users\<user>\AppData\Roaming)
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config directory: %w", err)
	}

	// Join the base config directory with your application's folder and then the file name.
	fullPath := filepath.Join(configDir, appConfigRelativePath, configFileName)
	return fullPath, nil
}

func main() {
	//var filePath = testFilePath
	filePath, err := getDefaultConfigPath()
	if err != nil {
		// TODO: figure out what to do with this err...
		return
	}

	// Define Command-Line Flags
	var quotesFilePathFlag string
	var quotesTagSearchFlag string
	var quotesAuthorSearchFlag string
	var versionFlag bool
	var quoteAdditionFlag bool

	// src file
	flag.StringVar(&quotesFilePathFlag, "file", filePath, "Path to the quotes file")
	flag.StringVar(&quotesFilePathFlag, "f", filePath, "Path to the quotes file")
	//tag search
	flag.StringVar(&quotesTagSearchFlag, "tag", "", "Tag to search quotes for (case-insensitive)")
	flag.StringVar(&quotesTagSearchFlag, "t", "", "Tag to search quotes for (case-insensitive)")
	// author search
	flag.StringVar(&quotesAuthorSearchFlag, "author", "", "Author to search quotes for (case-insensitive)")
	flag.StringVar(&quotesAuthorSearchFlag, "a", "", "Short for --author (case-insensitive)")
	// version
	flag.BoolVar(&versionFlag, "version", false, "Print application version")
	flag.BoolVar(&versionFlag, "v", false, "Print application version")
	// add new quote
	flag.BoolVar(&quoteAdditionFlag, "new", false, "Create new quote")
	flag.BoolVar(&quoteAdditionFlag, "n", false, "Create new quote")
	flag.Parse()

	// Display program version
	if versionFlag {
		fmt.Printf("Quote CLI Version: %s\n", appVersion)
		return
	}

	// Load Quotes
	quoteList, err := quotes.LoadQuotesFromFile(quotesFilePathFlag)
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}

	// quote addition
	if quoteAdditionFlag {
		display.DisplayQuoteAdditionPrompt(filePath)
		return
	}

	// quote searching
	if quotesTagSearchFlag != "" {
		// Handle tag search
		foundQuotes := quotes.SearchByQuoteTag(quoteList, quotesTagSearchFlag)
		if err != nil {
			log.Fatalf("Error with search: %v", err)
		}
		display.DisplayQuoteListWraped(foundQuotes)

	} else if quotesAuthorSearchFlag != "" {
		// Handle Author search
		foundQuotes := quotes.SearchByQuoteAuthor(quoteList, quotesAuthorSearchFlag)
		if err != nil {
			log.Fatalf("Error with search: %v", err)
		}
		display.DisplayQuoteListWraped(foundQuotes)

	} else {
		// Display Random Quote
		randomInt := rand.Intn(len(quoteList))
		display.DisplayQuoteWrapedBoarder(quoteList[randomInt])
		//display.DisplayQuoteWraped(quoteList[randomInt])
	}
}
