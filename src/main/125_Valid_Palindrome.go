package main

import (
	"fmt"
	"unicode"
)

type P125 struct {
}

func P125isPalindrome(s string) bool {
	var p P125
	return p.sol1(s)
}

// 这个函数返回值没有保存该字符串
func Reverse1(s []rune) []rune {
	neg := make([]rune, 0)
	for _, p := range s {
		defer func(v rune) {
			neg = append(neg, v)
		}(p)
	}

	return neg
}

// 这样处理可以成功返回反转的字符串，不过返回类型不是string类型
func Reverse_2(s string) *[]rune {
	var a []rune
	for _, k := range []rune(s) {
		defer func(v rune) {
			a = append(a, v)
		}(k)
	}
	return &a
}

func (p P125)sol1(s string) bool {
	pos := make([]rune, 0)
	for _, c := range s {
		if unicode.IsLetter(c) {
			pos = append(pos, unicode.ToLower(c))
		}
	}

	neg := Reverse1(pos)

	fmt.Println(string(pos))
	fmt.Println(string(neg))
	return string(pos) == string(neg)
}

func main()  {
	fmt.Println(P125isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(P125isPalindrome("race a car"))
}
