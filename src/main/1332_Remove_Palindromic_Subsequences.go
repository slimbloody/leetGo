package main

import "fmt"
/*
You need to know the difference between subarray and subsequence.
Subarray need to be consecutive。
Subsequence don't have to be consecutive.


If it's empty sting, return 0;
If it's palindrome, return 1;
Otherwise, we need at least 2 operation.
 */

type P1332 struct {
}

func Reverse(s string) string {
	runes := []rune(s)

	for i := 0; i < (len(s) - 1) / 2; i++ {
		runes[i], runes[len(s) - 1 - i] =  runes[len(s) - 1 - i], runes[i]
	}

	return string(runes)
}

func removePalindromeSub(s string) int {
	var p P1332
	return p.sol1(s)
}

func (p P1332)sol1(s string) int {
	if len(s) == 0 {
		return 0
	} else if Reverse(s) == s {
		return 1
	} else {
		return 2
	}
}

func main() {
	fmt.Println(removePalindromeSub("ababa"))
	fmt.Println(removePalindromeSub("abb"))
	fmt.Println(removePalindromeSub("baabb"))
	fmt.Println(removePalindromeSub(""))
}
