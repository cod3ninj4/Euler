package main

// The difference between the sum of the squares of the first 100 natural numbers and the square of the sum.

import (
	"fmt"
)

func square_sum(x int) int {
	var sum int = 0
	for i := 1; i <= x; i++ {
		sum = sum + i
	}
	return sum * sum
}

func square_natural(x int) int {
	var total int = 0
	for i := 1; i <= x; i++ {
		total = total + (i * i)
	}
	return total
}

func main() {
	fmt.Printf("Difference: %d\n", square_sum(100)-square_natural(100))
}
