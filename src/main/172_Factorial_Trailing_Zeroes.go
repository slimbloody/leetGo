package main

import "fmt"

type P172 struct {
}

func trailingZeroes(n int) int {
	var p P172
	return p.sol1(n)
}

func (p P172)sol1(n int) int {
	if n % 10 >= 5 {
		return n / 10 * 2 + 1
	} else {
		return n / 10 * 2
	}
}

func main()  {
	//fmt.Println(trailingZeroes(3))
	//fmt.Println(trailingZeroes(5))
	//fmt.Println(trailingZeroes(11))
	fmt.Println(trailingZeroes(30))
}

