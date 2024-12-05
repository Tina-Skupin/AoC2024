package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("01data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// create my lists
	var liste1 []int
	var liste2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split line into fields
		fields := strings.Fields(line)

		// We expect exactly 2 fields
		if len(fields) < 2 {
			fmt.Println("Error: line does not have 2 fields")
			continue
		}

		// Convert and append to Liste 1
		num1, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return
		}
		liste1 = append(liste1, num1)

		// convert and append to liste 2
		num2, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return
		}
		liste2 = append(liste2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Now `datasets` contains all the integer data

	sammel := 0.00

	sort.Ints(liste1)
	sort.Ints(liste2)

	for i := range liste1 {
		abstand := math.Abs(float64(liste1[i] - liste2[i]))
		sammel = sammel + abstand
	}
	fmt.Println(sammel)
}
