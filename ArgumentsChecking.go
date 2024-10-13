package main

import (
	"fmt"
	"os"
)

func ArgumentsChecking() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("usage: go run main.go <input_filename> <output_filename>")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	return ProcessFile(inputFile, outputFile)
}

func ProcessFile(inputFile, outputFile string) error {
	content, err := ReadInputFile(inputFile)
	if err != nil {
		return err
	}

	processedContent := CheckString(content)

	err = WriteOutputFile(outputFile, processedContent)
	if err != nil {
		return err
	}

	return nil
}

func WriteOutputFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}
	return nil
}

func ReadInputFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}
