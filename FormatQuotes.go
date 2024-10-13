package main

import (
	"regexp"
	"strings"
)

func FormatQuotes(text string) string {
	// Handle single quotes
	re := regexp.MustCompile(`'\s*([^']+?)\s*'`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		// Trim spaces inside the quotes
		inner := strings.TrimSpace(match[1 : len(match)-1])
		return "'" + inner + "'"
	})
}
