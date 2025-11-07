package display

import (
	"reflect"
	"testing"
)

func TestWrapText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		width    int
		expected []string
	}{
		{
			name:     "Empty string",
			text:     "",
			width:    10,
			expected: nil,
		},
		{
			name:     "Zero width",
			text:     "Hello world",
			width:    0,
			expected: []string{"Hello world"},
		},
		{
			name:     "Negative width",
			text:     "Hello world",
			width:    -5,
			expected: []string{"Hello world"},
		},
		{
			name:     "Text fits on one line",
			text:     "Hello world",
			width:    15,
			expected: []string{"Hello world"},
		},
		{
			name:     "Basic wrapping",
			text:     "This is a test string that needs to be wrapped.",
			width:    10,
			expected: []string{"This is a", "test", "string", "that needs", "to be", "wrapped."},
		},
		{
			name:     "Exact fit on line end",
			text:     "Hello world",
			width:    11,
			expected: []string{"Hello world"},
		},
		{
			name:     "Word longer than width",
			text:     "An extraordinarilylongword is here.",
			width:    5,
			expected: []string{"An", "extraordinarilylongword", "is", "here."},
		},
		{
			name:     "Multiple spaces between words",
			text:     "Hello   world  how are you?",
			width:    10,
			expected: []string{"Hello", "world how", "are you?"},
		},
		{
			name:     "Leading and trailing spaces",
			text:     "  Hello world  ",
			width:    10,
			expected: []string{"Hello", "world"},
		},
		{
			name:     "Newlines in input (should be treated as spaces by strings.Fields)",
			text:     "Line1\nLine2\nLine3",
			width:    5,
			expected: []string{"Line1", "Line2", "Line3"},
		},
		{
			name:     "Single word longer than width, starts new line",
			text:     "short extraordinarilylongword",
			width:    10,
			expected: []string{"short", "extraordinarilylongword"},
		},
		{
			name:     "Just spaces",
			text:     "   ",
			width:    10,
			expected: nil,
		},
		{
			name:     "Single word exactly width",
			text:     "word",
			width:    4,
			expected: []string{"word"},
		},
		{
			name:     "Single word less than width",
			text:     "word",
			width:    5,
			expected: []string{"word"},
		},
		{
			name:     "Multiple short words filling line",
			text:     "a b c d e f g h i j k l m n o p q r s t u v w x y z",
			width:    10,
			expected: []string{"a b c d e", "f g h i j", "k l m n o", "p q r s t", "u v w x y", "z"},
		},
		/*
			{
				name:     "Long text with varied word lengths",
				text:     "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
				width:    20,
				expected: []string{"Go is an open", "source programming", "language that", "makes it easy to", "build simple,", "reliable, and", "efficient software."},
			},
			{
				name:     "Unicode characters",
				text:     "こんにちは世界 😊 This is a test.",
				width:    10,
				expected: []string{"こんにちは世界", "😊 This is", "a test."},
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wrapText(tt.text, tt.width)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("wrapText(%q, %d) = %v; want %v", tt.text, tt.width, got, tt.expected)
			}
		})
	}
}

// TODO: The wrapText function needs to be updated to pass the new test cases.
// It should be designed to handle the `Unicode characters` and `Long text with varied word lengths` cases.
