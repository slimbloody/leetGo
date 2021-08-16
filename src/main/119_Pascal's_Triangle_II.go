package main

import "fmt"

type P119 struct {
}

func getRow(rowIndex int) []int {
	var p P119
	return p.sol5(rowIndex)
}

func (p P119) sol1(rowIndex int) []int {
	var res = make([]int, rowIndex + 1)
	res[0] = 1
	for i := 1; i < rowIndex + 1; i++ {
		for j := i; j >= 1; j-- {
			res[j] += res[j - 1]
		}
	}
	return res
}


/*
class Solution(object):
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        row = [1]
        for _ in range(rowIndex):
            row = [x + y for x, y in zip([0]+row, row+[0])]
        return row
 */

func (p P119) sol5(rowIndex int) []int {
	var res []int
	for i := 0; i <= rowIndex; i++ {
		var temp = make([]int, i + 1)
		temp[0] = 1
		for j := 1; j < i; j++ {
			temp[j] = res[j - 1] + res[j]
		}
		temp[len(temp) - 1] = 1
		res = temp
	}

	return res
}

func main()  {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}

