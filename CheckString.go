package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func CheckString(str string) string {
	re2 := regexp.MustCompile(`^[aeiouhAEIOUH]$`)
	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		if i+1 < len(words) && words[i+1] == "(hex)" {
			num, err := strconv.ParseInt(words[i], 16, 64)
			if err != nil {
				log.Printf("Error parsing hex: %v", err)
				continue
			}
			words[i] = strconv.Itoa(int(num))
			words = removeElement(words, i+1)
		} else if i+1 < len(words) && words[i+1] == "(bin)" {
			num, err := strconv.ParseInt(words[i], 2, 64)
			if err != nil {
				log.Printf("Error parsing binary: %v", err)
				continue
			}
			words[i] = strconv.Itoa(int(num))
			words = removeElement(words, i+1)
		} else if i+1 < len(words) && words[i+1] == "(up)" {
			words[i] = strings.ToUpper(words[i])
			words = removeElement(words, i+1)
		} else if i+1 < len(words) && words[i+1] == "(low)" {
			words[i] = strings.ToLower(words[i])
			words = removeElement(words, i+1)
		} else if i+1 < len(words) && words[i+1] == "(cap)" {
			words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
			words = removeElement(words, i+1)
		} else if words[i] == "a" || words[i] == "A" {
			if i+1 < len(words) {
				str := string(words[i+1][0])
				if re2.MatchString(str) {
					words[i] += "n"
				}
			}
		} else if i+1 < len(words) && strings.HasPrefix(words[i+1], "(") {
			modifierStr := words[i+1]
			removeCount := 1

			// Check if the modifier is split across two words
			if !strings.HasSuffix(modifierStr, ")") && i+2 < len(words) {
				modifierStr += " " + words[i+2]
				removeCount = 2
			}

			if isValidModifier(modifierStr) {
				inst, num, err := convertModifier(modifierStr)
				if err != nil {
					log.Printf("Error converting modifier: %v", err)
					continue
				}

				// Calculate the start index and ensure it's not out of bounds
				startIndex := i - num + 1
				if startIndex < 0 {
					startIndex = 0
				}

				// Apply transformations to the previous 'num' words, up to the startIndex
				for j := startIndex; j <= i; j++ {
					switch inst {
					case "low":
						words[j] = strings.ToLower(words[j])
					case "up":
						words[j] = strings.ToUpper(words[j])
					case "cap":
						words[j] = strings.Title(strings.ToLower(words[j]))
					}
				}

				// Remove the modifier word(s)
				words = append(words[:i+1], words[i+1+removeCount:]...)
				i-- // Adjust index after removal
			}
		}
	}

	formattedStr := strings.Join(words, " ")
	formattedStr = FormatPunctuations(formattedStr)
	formattedStr = FormatQuotes(formattedStr)
	return formattedStr
}
