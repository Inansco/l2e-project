package main

import (
	"fmt"
	"os"
)

func main() {
	// Check arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)

	// Call the main transformer (defined in transformations.go)
	text = Transform(text) + "\n"

	// Write output file
	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Done!")
}
