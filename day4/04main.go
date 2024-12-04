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

	// Call the other ufnction to find Vin Diesel
	xanderCount := findXander(formatted)
	fmt.Printf("Found 'Vin diesel' %d times.\n", xanderCount)
}

// Helper function to add quotes to each element in a string slice
func quote(row []string) []string {
	quotedRow := make([]string, len(row))
	for i, el := range row {
		quotedRow[i] = fmt.Sprintf("\"%s\"", el)
	}
	return quotedRow
}

func findXMAS(grid [][]string) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			// Check right
			if j <= len(grid[i])-4 && grid[i][j] == "X" && grid[i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S" {
				count++
			}
			// check backwards
			if j <= len(grid[i])-4 && grid[i][j] == "S" && grid[i][j+1] == "A" && grid[i][j+2] == "M" && grid[i][j+3] == "X" {
				count++
			}
			// Check down
			if i <= len(grid)-4 && grid[i][j] == "X" && grid[i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S"{
				count++
			}
			// Check down backwards
			if i <= len(grid)-4 && grid[i][j] == "S" && grid[i+1][j] == "A" && grid[i+2][j] == "M" && grid[i+3][j] == "X"{
				count++
			}
			// diagonal forward left to right
			if i <= len(grid)-4 && j <= len(grid[i])-4 && grid[i][j] == "X" && grid[i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S"{
				count++
			}
			// diagonal backward left to right
			if i <= len(grid)-4 && j <= len(grid[i])-4 && grid[i][j] == "S" && grid[i+1][j+1] == "A" && grid[i+2][j+2] == "M" && grid[i+3][j+3] == "X"{
				count++
			}
			// diagonal backward right to left
			if i <= len(grid)-4 && j <= len(grid[i])-4 && grid[i+3][j] == "S" && grid[i+2][j+1] == "A" && grid[i+1][j+2] == "M" && grid[i][j+3] == "X"{
				count++
			}
			// diagonal forward right to left
			if i <= len(grid)-4 && j <= len(grid[i])-4 && grid[i+3][j] == "X" && grid[i+2][j+1] == "M" && grid[i+1][j+2] == "A" && grid[i][j+3] == "S"{
				count++
			}

		}
	}
	return count
}

func findXander(grid [][]string) int {
	xanderCount := 0
	for i := range grid {
		for j := range grid[i] {
			// diagonal forward left to right
			if i <= len(grid)-3 && j <= len(grid[i])-3 && grid[i][j] == "M" && grid[i+1][j+1] == "A" && grid[i+2][j+2] == "S" && grid[i+2][j] == "M" && grid[i][j+2] == "S"{
				xanderCount++
			}
			if i <= len(grid)-3 && j <= len(grid[i])-3 && grid[i][j] == "M" && grid[i+1][j+1] == "A" && grid[i+2][j+2] == "S" && grid[i+2][j] == "S" && grid[i][j+2] == "M"{
				xanderCount++
			}
			if i <= len(grid)-3 && j <= len(grid[i])-3 && grid[i][j] == "S" && grid[i+1][j+1] == "A" && grid[i+2][j+2] == "M" && grid[i+2][j] == "M" && grid[i][j+2] == "S"{
				xanderCount++
			}
			if i <= len(grid)-3 && j <= len(grid[i])-3 && grid[i][j] == "S" && grid[i+1][j+1] == "A" && grid[i+2][j+2] == "M" && grid[i+2][j] == "S" && grid[i][j+2] == "M"{
				xanderCount++
			}
		}
	}
	return xanderCount
}
