package main
//Which starting number, under one million, produces the longest Collatz chain
import(
	"fmt"
)

func main() {
	largest := 0
	value := 0
	var num int64
	for x := 1000000; x >= 1; x-- {
		num = int64(x)
		count := 1
		for num > 1 {
			count++
			if num % 2 == 0 {
				num = num/2
			}else{
				num = ((3*num) + 1)
			}
		}
		if count > largest {
			largest = count
			value = x
		}
	}
	fmt.Println("Largest Sequence Using: ", value)
	fmt.Println("Sequence was: ", largest)
}
