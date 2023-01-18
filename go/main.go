package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 1
	for i := 0; i < 5; i++ {
		sum *= 10
	}
	fmt.Println(sum * (sum - 2))
	fmt.Println(math.MaxInt32)
}
