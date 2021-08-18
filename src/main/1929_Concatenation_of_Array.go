package main

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func main()  {
	fmt.Println(getConcatenation([]int{}))
	fmt.Println(getConcatenation([]int{1, 2, 3}))
	fmt.Println(getConcatenation([]int{1, 2, 3, 4}))
}