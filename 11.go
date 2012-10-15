package main
// find the largest product of any set of five consecutive numbers in the txt file (20x20 grid)
import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	var number [20][20]int
	var largest int64 = 0
	var cval int64 = 0
	content, err := ioutil.ReadFile("11.txt")
	if err != nil {
		fmt.Print("Error in Read\n")
	}
	lines := strings.Split(string(content), "\n")
	for x := 0; len(lines)> x+1; x++ {
		values := strings.Split(lines[x], " ")
		for y := 0; y < 20; y++ {
			number[x][y], err = strconv.Atoi(values[y])
			if err != nil {
				fmt.Print("Error in Atoi\n")
			}
		}
	}
	// Calculate across
	for x := 0; x < 20; x++ {
		for y := 0; y < 17; y++ {
			a := number[x][y]
			b := number[x][y+1]
			c := number[x][y+2]
			d := number[x][y+3]
			cval = int64(a * b * c * d)
			if cval > largest {
				largest = cval
			}
		}
	}
	// Calculate down
	for x := 0; x < 17; x++ {
		for y := 0; y < 20; y++ {
			a := number[x][y]
			b := number[x+1][y]
			c := number[x+2][y]
			d := number[x+3][y]
			cval = int64(a * b * c * d)
			if cval > largest {
				largest = cval
			}
		}
	}
	//Calculate down/right diag
	for x := 0; x < 17; x++ {
		for y := 0; y < 17; y++ {
			a := number[x][y]
			b := number[x+1][y+1]
			c := number[x+2][y+2]
			d := number[x+3][y+3]
			cval = int64(a * b * c * d)
			if cval > largest {
				largest = cval
			}
		}
	}
	//Calculate down/left diag
	for x := 0; x < 17; x++ {
		for y := 3; y < 20; y++ {
			a := number[x][y]
			b := number[x+1][y-1]
			c := number[x+2][y-2]
			d := number[x+3][y-3]
			cval = int64(a * b * c * d)
			if cval > largest {
				largest = cval
			}
		}
	}
	fmt.Printf("Largest Product: %d\n", largest)
}
