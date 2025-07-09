package display

import (
	"os"
	"fmt"
	"strings"
	"golang.org/x/term"
	"unicode/utf8"

	"quote-cli/internal/quotes"
)

// ====================================================== \\
//	Internal Helper Functions
// ====================================================== \\

// getTerminalWidth gets and returns the terminal width, if get
// fails falls back to 80 rune width.
func getTerminalWidth() int {
	fallbackWidth := 80
	fileDescriptor := int(os.Stdout.Fd()) // Get the file descriptor for standard output

	// Check if os.Stdout is actually connected to a terminal
	if term.IsTerminal(fileDescriptor) {
		width, _, err := term.GetSize(fileDescriptor) // GetSize returns width, height, error
		if err != nil {
			return fallbackWidth
		}

		if width > 0 {
			return width
		}
	}

	return fallbackWidth
}

// wrapText wraps the given text to the specified width, ensuring words are not broken.
func wrapText(text string, width int) string {
	if width <= 0 {
		return text
	}

	var wrappedLines []string
	currentLine := ""
	words := strings.Fields(text)

	for _, word := range words {
		// Calculate the length of the word in runes (Unicode characters)
		wordLen := utf8.RuneCountInString(word)

		// Calculate the length of the current line with the new word and a space
		currentLineLen := utf8.RuneCountInString(currentLine)

		// If adding the word (and a space if not the first word on the line)
		// would exceed the width, start a new line.
		// Also handle words longer than the line width (they'll occupy their own line)
		if currentLineLen+1+wordLen > width && currentLineLen > 0 {
			wrappedLines = append(wrappedLines, currentLine)
			currentLine = word
		} else if currentLineLen == 0 { // First word on a new line
			currentLine = word
		} else { // Add word to current line with a space
			currentLine += " " + word
		}

		// TODO: Handle cases where a single word is longer than the line width.
		// This strategy splits the word, but the prompt asked not to break words.
		// If a word is inherently too long for a single line, you have a choice:
		// 1. Let it exceed the width (simple, potentially ugly)
		// 2. Break it (not what was asked)
		// 3. Or just wrap it on its own line as done above.
		// The current logic places it on its own line if the current line already has content.
		// If the *word itself* is longer than 'width', it will still appear on one line
		// potentially exceeding the 'width'. To force a break, you'd need more logic.
		// The prompt specified "not broken up", so we let long words stand.
	}

	// Add the last line if it has content
	if currentLine != "" {
		wrappedLines = append(wrappedLines, currentLine)
	}

	return strings.Join(wrappedLines, "\n")
}

// ====================================================== \\
//	std Out Display Functions
// ====================================================== \\

// displayQuoteList prints a list of quotes to the console no fancy formatting.
func DisplayQuoteListWraped(quoteList []quotes.Quote) {
	for _, quote := range quoteList {
		DisplayQuoteWraped(quote)
	}
}

// displayQuote prints the quote to the console no fancy formatting.
func DisplayQuoteSimple(quote quotes.Quote) {
	fmt.Println("")
	fmt.Printf("“%s”\n", quote.Text)
	fmt.Printf("  - %s\n", quote.Author)
	fmt.Println("")
}

// DisplayQuoteWraped prints the quote wrapped to the console width.
func DisplayQuoteWraped(quote quotes.Quote) {
	terminalWidth := getTerminalWidth()

	// prep then print the quote
	fmt.Println("")
	wrappedQuote := wrapText(quote.Text, terminalWidth-4) // Subtract a bit for padding/border
	fmt.Printf("“%s”\n", wrappedQuote)
	fmt.Printf("  - %s\n", quote.Author)
	fmt.Println("")
}
