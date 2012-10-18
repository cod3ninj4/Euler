package main

import (
	"fmt"
)

func main() {
	number := 0
	for x := 1; x > 0 ; x++ {
		number += x
		var factors int = 0
		if number % 2 == 0 {
			for i := number; i > 0; i-- {
				if number % i == 0 {
					factors++
				}
			}
		}
		if factors >= 500 {
			fmt.Printf("Number: %d Factors: %d\n", number, factors)
			goto Done
		}
	}
	Done:
	
}
