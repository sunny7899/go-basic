package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Function to format JSON data
func formatJSON(input []byte) (string, error) {
	var formattedJSON bytes.Buffer
	err := json.Indent(&formattedJSON, input, "", "  ")
	if err != nil {
		return "", err
	}
	return formattedJSON.String(), nil
}

func main() {
	// Read JSON input from standard input
	fmt.Println("Enter JSON data:")
	var input []byte
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Format JSON
	output, err := formatJSON(input)
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		os.Exit(1)
	}

	// Print formatted JSON
	fmt.Println("Formatted JSON:")
	fmt.Println(output)
}