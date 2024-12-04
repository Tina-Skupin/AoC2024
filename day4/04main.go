package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Assume here that your file is named "data.txt"
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var formatted [][]string

	for scanner.Scan() {
		// Read each line from the file
		line := scanner.Text()
		row := strings.Split(line, "")
		formatted = append(formatted, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print the formatted matrix
	fmt.Println("Formatted Matrix:")
	for _, row := range formatted {
		fmt.Printf("{%s},\n", strings.Join(quote(row), ", "))
	}

	// Call the function to find "xmas"
	count := findXMAS(formatted)
	fmt.Printf("Found 'xmas' %d times.\n", count)
}

// Helper function to add quotes to each element in a string slice
func quote(row []string) []string {
	quotedRow := make([]string, len(row))
	for i, el := range row {
		quotedRow[i] = fmt.Sprintf("\"%s\"", el)
	}
	return quotedRow
}

// Place your findXMAS function here
func findXMAS(grid [][]string) int {
	// Implement your logic to find the word "xmas"
	return 0 // replace this with your searching logic
}
