package main

// Sum of all primes under 2 million
import (
	"fmt"
)

func is_prime(n int64) bool {
	var prime bool = false
	if n <= 1 {
		return prime
	}
	if n == 2 || n == 3 {
		prime = true
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

func main() {
	var sum int64 = 0
	var x int64 = 0
	for x = 1; x < 2000000; x++ {
		if is_prime(x) {
			sum = sum + x
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
