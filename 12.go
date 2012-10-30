package main

/*
  First triagle number with over 500 factors.
  Note, you only need to check up to the SQRT of the number b/c every
  factor below that has a mirror, hence factors + 2
*/
import (
	"fmt"
	"math"
)

func main() {
	number := 0
	for x := 1; x > 0; x++ {
		number += x
		factors := 0
		for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
			if number%i == 0 {
				factors += 2
			}
		}
		if factors > 500 {
			fmt.Printf("Number: %d Factors: %d\n", number, factors)
			goto Done
		}
	}
Done:
}
