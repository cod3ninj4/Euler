package main

import (
	"fmt"
)

func is_prime(n int64) bool {
	var prime bool = false
	if n <= 1 {
		return prime
	}
	if n%2 == 0 {
		return prime
	}
	var max int64 = n
	var x int64
	for x = 3; x < max+1; x = x + 2 {
		max = n / x
		if n%x == 0 {
			return prime
		}
	}
	prime = true
	return prime
}

func largest_factor(n int64) int64 {
	var largest int64 = 0
	var x int64
	for x = 3; x < n; x += 2 {
		largest = n / x
		if n%x == 0 {
			if is_prime(largest) {
				return largest
			}
		}
	}
	return largest
}

func main() {
	fmt.Printf("Largest Prime Factor: %d\n", largest_factor(600851475143))
}
