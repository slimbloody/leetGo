package main

import (
	"fmt"
)

type P74 struct {
}

// mysql 找页空间
func searchMatrix(matrix [][]int, target int) bool {
	var p P74
	return p.sol1(matrix, target)
}

func (p P74)sol1(matrix [][]int, target int) bool {
	var row = len(matrix) - 1
	var col int
	if row < 0 && matrix[row] != nil {
		return false
	} else {
		col = len(matrix[0]) - 1

		if col < 0 {
			return false
		}
	}

	var rLow, rHigh = 0, row
	var cLow, cHigh = 0, col
	var r int
	for rLow < rHigh - 1 {
		var rMid = rLow + (rHigh - rLow) >> 2
		if target > matrix[rMid][0] {
			rLow = rMid + 1
		} else if target < matrix[rMid][0] {
			rHigh = rMid
		} else {
			return true
		}
	}

	if target < matrix[rHigh][0] {
		r = rHigh
	} else {
		r = rLow
	}

	for cLow <= cHigh {
		var cMid = cLow + (cHigh - cLow) >> 2
		if target < matrix[r][cMid] {
			cHigh = cMid - 1
		} else if target > matrix[r][cMid] {
			cLow = cMid + 1
		} else if target == matrix[r][cMid] {
			return true
		}
	}

	return false
}

func main()  {
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3))
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{}}, 1))
	fmt.Println(searchMatrix([][]int{{1}, {3}}, 0))
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,50}}, 11))
}









