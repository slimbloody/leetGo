package main

import (
	"fmt"
	"sort"
	"strings"
)

type P1859 struct {
}

func sortSentence(s string) string {
	var p P1859
	return p.sol1(s)
}

func (p P1859) sol1(s string) string {
	var strMap = make(map[int]string)
	var key []int
	var resStrs []string
	strs := strings.Split(s, " ")
	for _, str := range strs {
		var k = int(str[len(str) - 1] - '0')
		key = append(key, k)
		strMap[k] = str[: len(str) - 1]
	}

	sort.Ints(key)

	for _, k := range key {
		resStrs = append(resStrs, strMap[k])
	}

	res := strings.Join(resStrs, " ")
	return res
}

func main()  {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}