package main

// Find the number of unique down/right routes through a 20x20 grid
import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	if n <= 0 {
		return big.NewInt(1)
	}
	ans := big.NewInt(0)
	return ans.Mul(big.NewInt(n), factorial(n-1))
}

/*
  the formula for number of routes is 
  (n!)/(x! * y!) where n = x + y
  n possible steps with x steps right and y steps down
*/
func findRoutes(x int64, y int64) *big.Int {
	routes := big.NewInt(0)
	routes = routes.Mul(factorial(x), factorial(y))
	return routes.Div(factorial(x+y), routes)
}
func main() {
	routes := findRoutes(20, 20)
	fmt.Printf("Routes: %v\n", routes)
}
