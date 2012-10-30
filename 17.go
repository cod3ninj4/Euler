package main

import (
	"fmt"
)

func main() {
	total := 0
//Numbers 1-9 simple counts
	total = total + (3 + 3 + 5 + 4 + 4 + 3 + 5 + 5 + 4)
	oneToNine := total
//Numbers 10-19 simple counts
	total = total + (3 + 6 + 6 + 8 + 8 + 7 + 7 + 9 + 8 + 8)
//Numbers 20-99
	//10 times each prefix "Twentfy, Thirty..."
	total = total + (10 * (6 + 6 + 5 + 5 + 5 + 7 + 6 + 6))
	//Numbers 1-9 are repeated for each
	total = total + (8 * (3 + 3 + 5 + 4 + 4 + 3 + 5 + 5 + 4))
	oneToNinetyNine := total
//Numbers 100-999
	//1-9 Happens 100 times for each one hundred ...
	total = total + (oneToNine * 100)
	//1-99 hapens 9 times
	total = total + (oneToNinetyNine * 9)
	//The word "hundred" happens 9 times
	total = total + (7 * 9)
	//The words "hundred and" happen 99 times each 100 so 9 times..
	total = total + (10 * 99 * 9)
//1000
	total = total + 11

	fmt.Println("Total: ",total)
}
