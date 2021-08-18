package main

import "fmt"

type P1720 struct {
}

/*
A XOR B XOR B = A XOR (B XOR B) = A XOR 0 = A
 */
func decode(encoded []int, first int) []int {
	var p P1720
	return p.sol1(encoded, first)
}

func (p P1720)sol1(encoded []int, first int) []int {
	var res = make([]int, len(encoded) + 1)
	res[0] = first

	for i, num := range encoded {
		res[i + 1] = num ^ res[i]
	}
	
	return res
}

func main()  {
	// 1,0,2,1
	fmt.Println(decode([]int{1,2,3}, 1))
	// 4,2,0,7,4
	fmt.Println(decode([]int{6,2,7,3}, 4))
}

