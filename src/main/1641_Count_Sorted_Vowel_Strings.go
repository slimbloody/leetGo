package main

import "fmt"

type P1641 struct {
}

func countVowelStrings(n int) int {
	var p P1641
	return p.sol1(n)
}

func (p P1641)sol1(n int) int {
}

func main()  {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
}
