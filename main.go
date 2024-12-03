package main

import (
	"fmt"
	"math"
)

func innerLogic(dataset []int) int {
	safeReports := 0
	isRising, isFalling, isDistance := true, true, true

	for i := 1; i < len(dataset); i++ {
		if dataset[i-1] >= dataset[i] {
			isRising = false
		} 
		if dataset[i-1] <= dataset[i] {
			isFalling = false
		}
		distance := math.Abs(float64(dataset[i] - dataset[i-1]))
		if distance <1 || distance > 3{
			isDistance = false
		}
	}
	if isRising || isFalling && isDistance{
		safeReports++
	}
	return safeReports
}

func main() {
	// define dataset
	dataset := []int{7, 6, 5, 2, 1}

	//call the function
	result := innerLogic(dataset)

	fmt.Println(result)
}
