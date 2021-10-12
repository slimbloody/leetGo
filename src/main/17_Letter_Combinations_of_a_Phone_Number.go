package main

import "fmt"

type P17 struct {
}

func letterCombinations(digits string) []string {
	var p P17
	return p.sol1(digits)
}

func (p P17)sol1(digits string) []string {
	res := make([]string, 0)
	var keyboard []string = []string{
		"", "abc", "def", "ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz",
	}

	for _, char := range digits {
		num := char - '1'
		if len(res) == 0 {
			for _, key := range keyboard[num] {
				res = append(res, string(key))
			}
		} else {
			newRes := make([]string, 0)

			for _, key := range keyboard[num] {
				for _, r := range res {
					newRes = append(newRes, r + string(key))
				}
			}

			res = newRes
		}
	}

	return res
}

func main ()  {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
}