package main

import (
	"fmt"
)

type P1817 struct {
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	var p P1817
	return p.sol1(logs, k)
}

func (p P1817)sol1(logs [][]int, k int) []int {
	var resMap = make(map[int]map[int]int)
	for _, log := range logs {
		if _, firstOk := resMap[log[0]]; !firstOk {
			resMap[log[0]] = make(map[int]int)
		}
		resMap[log[0]][log[1]] = 1
	}

	var res = make([]int, k)
	for _, v := range resMap {
		res[len(v) - 1]++
	}

	return res
}

func main()  {
	fmt.Println(findingUsersActiveMinutes([][]int{{0,5},{1,2},{0,2},{0,5},{1,3}}, 5))
	fmt.Println(findingUsersActiveMinutes([][]int{{1,1},{2,2},{2,3}}, 4))
}
