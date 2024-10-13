package main

import "strings"

func FormatPunctuations(text string) string {
	// Handle special punctuation groups
	specialGroups := []string{"...", "!?"}
	for _, group := range specialGroups {
		text = strings.ReplaceAll(text, " "+group, group)
		text = strings.ReplaceAll(text, group+" ", group+" ")
	}

	// Handle single punctuation marks
	singlePunct := []string{".", ",", "!", "?", ":", ";"}
	for _, punct := range singlePunct {
		// Remove space before the punctuation
		text = strings.ReplaceAll(text, " "+punct, punct)

		// Ensure there's a space after the punctuation if it's not at the end of the text
		// and not followed by another punctuation mark
		for i := 0; i < len(text)-1; i++ {
			if text[i] == punct[0] && text[i+1] != ' ' {
				isNextPunct := false
				for _, p := range singlePunct {
					if text[i+1] == p[0] {
						isNextPunct = true
						break
					}
				}
				if !isNextPunct {
					text = text[:i+1] + " " + text[i+1:]
				}
			}
		}
	}

	return text
}
