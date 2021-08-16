package main

import "fmt"

type P1332 struct {
}

func getKth(lo int, hi int, k int) int {
	var p P1332
	return p.sol1(lo, hi, k)
}

func (p P1332) sol1(lo int, hi int, k int) int {
}

func main()  {
	fmt.Println(getKth(12, 15, 2))
	fmt.Println(getKth(1, 1, 1))
	fmt.Println(getKth(7, 11, 4))
	fmt.Println(getKth(10, 20, 5))
	fmt.Println(getKth(1, 1000, 777))
}
