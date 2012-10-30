package main

// Smallest positive number that is evenly divisible by number 1 - 20

import (
	"fmt"
)

func is_divisible(x, y int) bool {
	var result bool = false
	if x%y == 0 {
		result = true
	}
	return result
}

func divis_1_to_20(x int) bool {
	var divis bool = true
	for divis == true {
		for i := 3; i <= 20; i += 1 {
			divis = is_divisible(x, i)
			if divis == false {
				i = 99
			}
		}
		return divis
	}
	return divis
}

func main() {
	var value int = 2518
	var divisible bool = false
	for divisible != true {
		value += 2
		divisible = divis_1_to_20(value)
	}

	fmt.Printf("Smallest value is: %d\n", value)

}
