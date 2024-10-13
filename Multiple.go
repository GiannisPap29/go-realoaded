package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidModifier(str string) bool {
	// Remove outer parentheses
	str = strings.Trim(str, "()")

	// Split into parts
	parts := strings.SplitN(str, ",", 2)
	if len(parts) != 2 {
		return false
	}

	// Check the first part (word)
	word := strings.TrimSpace(parts[0])
	if word != "up" && word != "cap" && word != "low" {
		return false
	}

	// Check the second part (number)
	numStr := strings.TrimSpace(parts[1])
	_, err := strconv.Atoi(numStr)
	return err == nil
}

func convertModifier(str string) (string, int, error) {
	// Remove outer parentheses
	str = strings.Trim(str, "()")

	// Split into parts
	parts := strings.SplitN(str, ",", 2)
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid format")
	}

	// Extract and validate word
	word := strings.TrimSpace(parts[0])
	if word != "up" && word != "cap" && word != "low" {
		return "", 0, fmt.Errorf("invalid modifier: %s", word)
	}

	// Extract and validate number
	numStr := strings.TrimSpace(parts[1])
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid number: %v", err)
	}

	return word, num, nil
}
