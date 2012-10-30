package main

// Sum of all numbers less than 1000 that are divisible by 3 and 5
import (
	"fmt"
	"math"
)

func div_by_3(x float64) bool {
	var result bool = false
	if math.Remainder(x, 3) == 0 {
		result = true
	}
	return result
}
func div_by_5(x float64) bool {
	var result bool = false
	if math.Remainder(x, 5) == 0 {
		result = true
	}
	return result
}

func main() {
	var sum int = 0
	for x := 0; x <= 1000; x++ {
		if div_by_3(float64(x)) {
			sum = sum + x
		} else if div_by_5(float64(x)) {
			sum = sum + x
		}
	}
	fmt.Printf("Sum is: %d\n", sum)
}
