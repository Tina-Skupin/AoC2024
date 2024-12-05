package main

import (
	"fmt"
	"sort"
	"math"
)



func main() {
	liste1 := [] int{3,4,2,1,3,3}
	liste2 := [] int{4,3,5,3,9,3}

	sammel := 0.00

	sort.Ints(liste1)
	sort.Ints(liste2)

	for i := range liste1 {
		abstand := math.Abs(float64(liste1[i] - liste2[i]))
		fmt.Println(abstand)
		sammel = sammel + abstand
		}
		fmt.Println(sammel)
	}


