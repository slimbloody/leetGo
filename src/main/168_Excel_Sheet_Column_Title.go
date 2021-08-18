package main

import (
	"fmt"
	"strings"
)

func convertToTitle(columnNumber int) string {
	var sb strings.Builder
	var alphabeticSize = 26
	for remain := columnNumber; remain > 0; remain = (remain - 1) / alphabeticSize {
		sb.WriteString(string(rune(((remain - 1) % 26) + int('A'))))
	}
	r := []rune(sb.String())
	var returnRune []rune
	for i := len(r) - 1; i >=0; i-- {
		returnRune = append(returnRune, r[i])
	}

	return string(returnRune)
}

func main() {
	//res1 := convertToTitle(1)
	//res2 := convertToTitle(10)
	//res3 := convertToTitle(27)
	//res4 := convertToTitle(28)
	// ZY
	//res5 := convertToTitle(701)
	// AHP
	//res6 := convertToTitle(900)
	// FXSHRXW
	//res7 := convertToTitle(2147483647)
	res8 := convertToTitle(26)
	//fmt.Println(res1)
	//fmt.Println(res2)
	//fmt.Println(res3)
	//fmt.Println(res4)
	//fmt.Println(res5)
	//fmt.Println(res6)
	//fmt.Println(res7)
	fmt.Println(res8)
}