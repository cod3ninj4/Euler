package main
// There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find the product abc.

import (
	"fmt"
)

func main(){
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++{
			for c := b +1; c < 1000; c++{
				if (a * a) + (b * b) == (c * c) {
					if a + b + c == 1000 {
						fmt.Printf("A: %d B: %d C: %d\n", a, b, c)
						fmt.Printf("Product: %d\n", a * b * c)
						goto Done
					}
				}
			}
		}
	}	
Done:
}
