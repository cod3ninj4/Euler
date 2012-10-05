package main

import (
	"fmt"
)

func is_prime(n int64) bool{
	var prime bool = false
	if n <= 1 {
		return prime
	}
	if n % 2 == 0{
		return prime
	}
	var max int64 = n
	var x int64
	for x = 3; x < max +1; x = x + 2 {
		max = n / x
		if n % x == 0 {
			return prime
		}
	}
	prime = true
	return prime
}

func find_prime(x int) int{
	var num int = 0
	var value int = 0
	var place int = 2
	for place < x {
		if is_prime(int64(num)) {
			value = num
			place++
		}
		num++
	}
	return value
}

func main () {
	fmt.Printf("10,001 Prime: %d\n", find_prime(10001))	
}
