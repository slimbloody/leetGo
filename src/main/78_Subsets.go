package main

import (
	"fmt"
	"sort"
)

type P78 struct {
}

func subsets(nums []int) [][]int {
	var p P78
	return p.sol1(nums)
}

func (p P78)sol1(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}

	for _, num := range nums {
		curLen := len(res)
		for i := 0; i < curLen; i++ {
			addArr := append(res[i], num)
			res = append(res, addArr)
		}
	}

	return res
}

func main() {
	fmt.Println(subsets([]int{9,0,3,5,7}))
	//fmt.Println(subsets([]int{1,2,3}))
	//fmt.Println(subsets([]int{0}))
	//fmt.Println(subsets([]int{}))
}