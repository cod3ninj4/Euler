package main

// Find the largest palindrome made from the product of two 3-digit numbers.

import (
	"fmt"
	"strconv"
)

func is_palindrome(s string) bool {
	return reverse(s) == s
}

func reverse(s string) string {
	orig := string(s)
	len := len(s)
	n := make([]byte, len)
	for i := 0; i < len; i++ {
		n[len-i-1] = orig[i]
	}
	return string(n)
}

func main() {
	var largest int = 0
	for i := 999; i > 99; i-- {
		for x := 999; x > 99; x-- {
			if is_palindrome(strconv.Itoa(x * i)) {
				if i*x > largest {
					largest = i * x
				}
			}
		}
	}
	fmt.Printf("Largest Palindrome: %d\n", largest)
}
