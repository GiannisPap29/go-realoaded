package main

import (
	"testing"
)

func TestCheckString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Hex Conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Binary Conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Uppercase Conversion",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Lowercase Conversion",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Capitalize Conversion",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			name:     "Multiple Word Conversion",
			input:    "This is so exciting (up, 2)",
			expected: "this is SO EXCITING",
		},
		{
			name:     "A to An Conversion",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		{
			name:     "Complex Sentence",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "Empty String",
			input:    "",
			expected: "",
		},
		{
			name:     "Multiple Consecutive Conversions",
			input:    "test (up) (low) (cap) result",
			expected: "TEST result",
		},
		{
			name:     "A/An with H",
			input:    "A honest man with a hour to spare",
			expected: "An honest man with an hour to spare",
		},
		{
			name:     "Invalid Conversions",
			input:    "Test (invalid) and (up, invalid) and FF (hex) and 1010 (bin)",
			expected: "Test (invalid) and (up, invalid) and 255 and 10",
		},
		{
			name:     "Modifier at start",
			input:    "(up, 2) This should be uppercase",
			expected: "THIS SHOULD be uppercase",
		},
		{
			name:     "Modifier at end",
			input:    "This should not change (up, 2)",
			expected: "This should not change",
		},
		{
			name:     "Overlapping modifiers",
			input:    "one two three (up, 2) four (low, 2) five",
			expected: "ONE TWO three four five",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckString(tt.input)
			if result != tt.expected {
				t.Errorf("CheckString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatPunctuations(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic punctuation", "Hello , world !", "Hello, world!"},
		{"Ellipsis", "Thinking ... done.", "Thinking... done."},
		{"Double punctuation", "Really !?", "Really!?"},
		{"Mixed punctuation", "Hi ; there : how are you ?", "Hi; there: how are you?"},
		{"No spaces to remove", "Already,formatted.text", "Already,formatted.text"},
		{"Multiple spaces", "Too   many    spaces", "Too   many    spaces"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatPunctuations(tt.input)
			if result != tt.expected {
				t.Errorf("FormatPunctuations() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic quotes", "He said ' hello '", "He said 'hello'"},
		{"Multiple words", "' Testing 1 2 3 '", "'Testing 1 2 3'"},
		{"No quotes", "No quotes here", "No quotes here"},
		{"Unmatched quotes", "' Unmatched", "' Unmatched"},
		{"Multiple quote pairs", "' First ' and ' Second '", "'First' and 'Second'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatQuotes(tt.input)
			if result != tt.expected {
				t.Errorf("FormatQuotes() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsValidModifier(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"(up, 3)", true},
		{"(low, 2)", true},
		{"(cap, 1)", true},
		{"(up, 0)", true},
		{"(invalid, 3)", false},
		{"(up, invalid)", false},
		{"(up,3)", false},         // No space after comma
		{"up, 3", false},          // Missing parentheses
		{"(up, 3, extra)", false}, // Too many parts
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isValidModifier(tt.input)
			if result != tt.expected {
				t.Errorf("isValidModifier(%s) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertModifier(t *testing.T) {
	tests := []struct {
		input        string
		expectedInst string
		expectedNum  int
		expectError  bool
	}{
		{"(up, 3)", "up", 3, false},
		{"(low, 2)", "low", 2, false},
		{"(cap, 1)", "cap", 1, false},
		{"(up, 0)", "up", 0, false},
		{"(invalid, 3)", "", 0, true},
		{"(up, invalid)", "", 0, true},
		{"(up,3)", "", 0, true}, // No space after comma
		{"up, 3", "", 0, true},  // Missing parentheses
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			inst, num, err := convertModifier(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("convertModifier(%s) error = %v, expectError %v", tt.input, err, tt.expectError)
				return
			}
			if inst != tt.expectedInst {
				t.Errorf("convertModifier(%s) inst = %v, want %v", tt.input, inst, tt.expectedInst)
			}
			if num != tt.expectedNum {
				t.Errorf("convertModifier(%s) num = %v, want %v", tt.input, num, tt.expectedNum)
			}
		})
	}
}
