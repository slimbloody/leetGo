package main

import (
	"fmt"
	"math"
)

type P338 struct {
}

func countBits(n int) []int {
	var p P338
	return p.sol1(n)
}

func (p P338)sol1(n int) []int {
	length := (n - 1) ^ n
	fmt.Println(length)
	res := make([]int, length)

	for i := 0; i < length && n != 0; i++ {
		if n - int(math.Exp2(float64(length-i-1))) > 0 {
			res[i] = 1
		}
	}

	return res
}

func main()  {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}

