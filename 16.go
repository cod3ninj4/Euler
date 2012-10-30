package main

//Sum of all the digits in 2^1000
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	line := strings.Split(strconv.FormatFloat(math.Pow(2, 1000), 'f', 0, 64), "")
	total := 0
	for x := 0; len(line) > x; x++ {
		holder, _ := strconv.Atoi(line[x])
		total = total + holder
	}
	fmt.Println(total)
}
