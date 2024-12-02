package main

import (
	"fmt"
)

func innerLogic(dataset []int) int {
	safeReports := 0
	isRising, isFalling := true, true

	for i := 1; i < len(dataset); i++ {
		if dataset[i-1] >= dataset[i] {
			isRising = false
		} else if dataset[i-1] <= dataset[i] {
			isFalling = false
		}
	}
	if isRising || isFalling {
		safeReports++
	}
	return safeReports
}

func main() {
	// define dataset
	dataset := []int{7, 6, 4, 2, 1}

	//call the function
	result := innerLogic(dataset)

	fmt.Println(result)
}
