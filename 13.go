package main

//Find the first 10 digits of the sum of 13.txt
import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

func main() {
	total := big.NewInt(0)
	holder := big.NewInt(0)
	content, _ := ioutil.ReadFile("13.txt")
	lines := strings.Split(string(content), "\n")
	for x := 0; len(lines) > x+1; x++ {
		holder, _ = holder.SetString(lines[x], 10)
		total.Add(total, holder)
	}
	fmt.Println("Total: ", total.String()[:10])
}
