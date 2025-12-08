package display

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"

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
// retruns an array of lines
func wrapText(text string, width int) []string {
	var wrappedLines []string
	currentLine := ""
	words := strings.Fields(text)

	if width <= 0 {
		return append(wrappedLines, text)
	}

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

	return wrappedLines
}

func basicWrapText(text string, width int) string {
	wrapedLines := wrapText("\t"+text, width)
	lineReturn := "\""

	for i, line := range wrapedLines {
		if i != 0 {
			lineReturn += "\n" + line
		} else {
			lineReturn += line
		}
	}
	lineReturn += "\""

	return lineReturn
}

func complexWrapText(text string, width int) string {
	wrapedLines := wrapText(text, width-10)
	returnLine := ""

	for i, line := range wrapedLines {
		fenceLine := ""

		// start fence
		if i == 0 {
			fenceLine += " |     " + line
		} else {
			fenceLine += " | " + line
		}

		// middle
		for len(fenceLine) < width-1 {
			fenceLine += " "
		}

		// end fence
		if i == len(wrapedLines)-1 {
			fenceLine += " |"
		} else {
			fenceLine += " |\n"
		}

		returnLine += fenceLine
	}

	return returnLine
}

// ====================================================== \\
//	std Out Display Functions
// ====================================================== \\

// child func of DisplayQuoteAdditionPrompt
func readQuote() string {
	newText := ""
	fmt.Print("Enter your quote: ")

	// read line
	reader := bufio.NewReader(os.Stdin)
	newText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading quote input:", err)
		return ""
	}
	newText = strings.TrimSuffix(newText, "\n")

	return newText
}

// child func of DisplayQuoteAdditionPrompt
func readAuthor() string {
	author := ""
	fmt.Print("Enter author name: ")

	// read line
	reader := bufio.NewReader(os.Stdin)
	author, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading author input:", err)
		return ""
	}
	author = strings.TrimSuffix(author, "\n")

	return author
}

// child func of DisplayQuoteAdditionPrompt
func readTags() []string {
	tags := []string{}
	noQuite := true
	for noQuite {
		newTag := ""
		fmt.Print("Enter quote tag (type Done to exit): ")

		// read line
		reader := bufio.NewReader(os.Stdin)
		newTag, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading quote input:", err)
		}
		newTag = strings.TrimSuffix(newTag, "\n")

		if strings.ToLower(newTag) == "done" {
			noQuite = false
		} else {
			tags = append(tags, newTag)
		}
	}

	return tags
}

// prompts for the new quote text, author and tags
func DisplayQuoteAdditionPrompt(filePath string) {
	newText := readQuote()
	if len(newText) <= 0 {
		fmt.Println("No new quote added")
		return
	}
	author := readAuthor()
	tags := readTags()

	// TODO: check for existing quote

	// TODO: if finds match allow exit or addition

	// add and catch err
	err := quotes.AddNewQuote(newText, author, tags, filePath)
	if err != nil {
		fmt.Println(err)
	}
}

// displayQuoteList prints a list of quotes to the console no fancy formatting.
func DisplayQuoteListWraped(quoteList []quotes.Quote) {
	for _, quote := range quoteList {
		DisplayQuoteWraped(quote)
	}
}

// displayQuote prints the quote to the console no fancy formatting.
func DisplayQuoteSimple(quote quotes.Quote) {
	fmt.Printf("%s\n", quote.Text)
	fmt.Printf("  - %s\n", quote.Author)
}

// DisplayQuoteWraped prints the quote wrapped to the console width.
func DisplayQuoteWraped(quote quotes.Quote) {
	terminalWidth := getTerminalWidth()

	// prep then print the quote
	wrappedQuote := basicWrapText(quote.Text, terminalWidth-4) // Subtract a bit for padding/border
	fmt.Printf("%s\n", wrappedQuote)
	fmt.Printf("  - %s\n", quote.Author)
}

// DisplayQuoteWraped prints the quote wrapped to the console width.
func DisplayQuoteWrapedBoarder(quote quotes.Quote) {
	terminalWidth := min(getTerminalWidth(), 90) // keep the quote/boarder from getting to long
	paddingMargin := 4
	wrappedQuote := complexWrapText(quote.Text, terminalWidth-paddingMargin) // Subtract a bit for padding/border

	capString := " "
	for range terminalWidth - paddingMargin {
		capString += "-"
	}

	author := quote.Author
	for len(author) < terminalWidth-paddingMargin-6 {
		author += " "
	}

	// prep then print the quote
	fmt.Printf("%s\n", capString)
	fmt.Printf("%s\n", wrappedQuote)
	fmt.Printf(" | - %s\n", author+" |")
	fmt.Printf("%s\n", capString)
}
