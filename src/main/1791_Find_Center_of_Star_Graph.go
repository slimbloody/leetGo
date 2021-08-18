package main

import "fmt"

type P1791 struct {
}

func findCenter(edges [][]int) int {
	var p P1791
	return p.sol1(edges)
}

func (p P1791)sol1(edges [][]int) int {
	var res int
	var count = 0
	var rMap = make(map[int]int)

	for _, arr := range edges {
		for _, num := range arr {
			if _, ok := rMap[num]; ok {
				rMap[num] += 1
			} else {
				rMap[num] = 1
			}
		}
	}

	for k, v := range rMap {
		if count < v {
			res = k
			count = v
		}
	}

	return res
}

func main()  {
	fmt.Println(findCenter([][]int{{1,2},{2,3},{4,2}}))
	fmt.Println(findCenter([][]int{{1,2},{5,1},{1,3},{1,4}}))
}
