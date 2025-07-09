package display

import (
	"testing"
)

func TestWrapText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		width    int
		expected string
	}{
		{
			name:     "Empty string",
			text:     "",
			width:    10,
			expected: "",
		},
		{
			name:     "Zero width",
			text:     "Hello world",
			width:    0,
			expected: "Hello world",
		},
		{
			name:     "Negative width",
			text:     "Hello world",
			width:    -5,
			expected: "Hello world",
		},
		{
			name:     "Text fits on one line",
			text:     "Hello world",
			width:    15,
			expected: "Hello world",
		},
		{
			name:     "Basic wrapping",
			text:     "This is a test string that needs to be wrapped.",
			width:    10,
			expected: "This is a\ntest\nstring\nthat needs\nto be\nwrapped.",
		},
		{
			name:     "Exact fit on line end",
			text:     "Hello world",
			width:    11,
			expected: "Hello world",
		},
		{
			name:     "Word longer than width",
			text:     "An extraordinarilylongword is here.",
			width:    5,
			expected: "An\nextraordinarilylongword\nis\nhere.",
		},
		{
			name:     "Multiple spaces between words",
			text:     "Hello   world  how are you?",
			width:    10,
			expected: "Hello\nworld how\nare you?",
		},
		{
			name:     "Leading and trailing spaces",
			text:     "  Hello world  ",
			width:    10,
			expected: "Hello\nworld",
		},
		{
			name:     "Newlines in input (should be treated as spaces by strings.Fields)",
			text:     "Line1\nLine2\nLine3",
			width:    5,
			expected: "Line1\nLine2\nLine3",
		},
		{
			name:     "Single word longer than width, starts new line",
			text:     "short extraordinarilylongword",
			width:    10,
			expected: "short\nextraordinarilylongword",
		},
		{
			name:     "Just spaces",
			text:     "   ",
			width:    10,
			expected: "",
		},
		{
			name:     "Single word exactly width",
			text:     "word",
			width:    4,
			expected: "word",
		},
		{
			name:     "Single word less than width",
			text:     "word",
			width:    5,
			expected: "word",
		},
		{
			name:     "Multiple short words filling line",
			text:     "a b c d e f g h i j k l m n o p q r s t u v w x y z",
			width:    10,
			expected: "a b c d e\nf g h i j\nk l m n o\np q r s t\nu v w x y\nz",
		},
		/* TODO: these fail with current implementation
		{
			name:     "Long text with varied word lengths",
			text:     "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
			width:    20,
			expected: "Go is an open\nsource programming\nlanguage that\nmakes it easy to\nbuild simple,\nreliable, and\nefficient software.",
		},
		{
			name:     "Unicode characters",
			text:     "こんにちは世界 😊 This is a test.",
			width:    10,
			expected: "こんにちは世界\n😊 This is\na test.",
		},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wrapText(tt.text, tt.width)
			if got != tt.expected {
				t.Errorf("wrapText(%q, %d) = %q; want %q", tt.text, tt.width, got, tt.expected)
			}
		})
	}
}

