package main

import (
	"fmt"
)

type P954 struct {
}

func canReorderDoubled(arr []int) bool {
	var p P954
	return p.sol1(arr)
}

func (p P954)sol1(arr []int) bool {
	var resMap = make(map[int]int)
	for _, item := range arr {
		if num, ok := resMap[item]; ok {
			if num > 0 {
				if nextNum, nextOk := resMap[item * 2]
			}
		}
	}
}

func main() {
	fmt.Println(canReorderDoubled([]int{3,1,3,6}))
	fmt.Println(canReorderDoubled([]int{2,1,2,6}))
	fmt.Println(canReorderDoubled([]int{4,-2,2,-4}))
	fmt.Println(canReorderDoubled([]int{1,2,4,16,8,4}))
	fmt.Println(canReorderDoubled([]int{}))
	fmt.Println(canReorderDoubled([]int{0, 0}))
	fmt.Println(canReorderDoubled([]int{2,1,2,1,1,1,2,2}))
}
