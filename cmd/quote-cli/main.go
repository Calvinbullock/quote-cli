package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"quote-cli/internal/quotes"
	"quote-cli/internal/display"
)

const appVersion = "1.0.0"

// path for testing
const defaultFilePath = "_assets/default.json"

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
	// var filePath = defaultFilePath
	filePath, err := getDefaultConfigPath()
	if err != nil {
		// TODO: figure out what to do with this err...
		return
	}

	// Define Command-Line Flags
	var quotesFilePath string
	var quotesTagSearch string
	var quotesAuthorSearch string
	var versionFlag bool

	// src file
	flag.StringVar(&quotesFilePath, "file", filePath, "Path to the quotes file")
	flag.StringVar(&quotesFilePath, "f", filePath, "Path to the quotes file")
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
		foundQuotes := quotes.SearchByQuoteTag(quoteList, quotesTagSearch)
		if err != nil {
			log.Fatalf("Error with search: %v", err)
		}
		display.DisplayQuoteListWraped(foundQuotes)

	} else if quotesAuthorSearch != "" {
		// Handle Author search
		foundQuotes := quotes.SearchByQuoteAuthor(quoteList, quotesAuthorSearch)
		if err != nil {
			log.Fatalf("Error with search: %v", err)
		}
		display.DisplayQuoteListWraped(foundQuotes)

	} else {
		// Display Random Quote
		randomInt := rand.Intn(len(quoteList))
		//display.DisplayQuoteSimple(quoteList[randomInt])
		display.DisplayQuoteWraped(quoteList[randomInt])
	}
}
