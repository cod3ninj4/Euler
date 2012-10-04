package main
// Find the sum of all even numbers in the Fibonacci sequence below 4mil
import(
	"fmt"
	"math"
)

func is_even(x float64) bool{
        var result bool = false
        if math.Remainder(x, 2) == 0 {
                result = true
        }
        return result
}



func main(){
	var x int = 1
	var y int = 2
	var sum int = 0

	for z := 0; z < 4000000; {
		if is_even(float64(y)){
			sum = sum + y
		}
		z = x + y
		x = y
		y = z
	}
	fmt.Printf("Sum: %d\n",sum)
}
