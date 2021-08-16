package main

import "fmt"

type P1646 struct {
}

func getMaximumGenerated(n int) int {
	var p P1646
	return p.sol1(n)
}

func (p P1646)sol1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		if n % 2 == 0 {
			return p.sol1(n / 2) * 2
		} else {
			return p.sol1(n / 2) * p.sol1(n / 2 + 1)
		}
	}
}

func main()  {
	fmt.Println(getMaximumGenerated(7))
	fmt.Println(getMaximumGenerated(2))
	fmt.Println(getMaximumGenerated(3))
}
