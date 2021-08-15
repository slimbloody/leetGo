package main

import (
	"fmt"
	"sort"
)

type P49 struct {
}

func groupAnagrams(strs []string) [][]string {
	var p P49
	return p.sol1(strs)
}

func (p P49)sol1(strs []string) [][]string {
	var res [][]string
	var resMap = make(map[string][]string)

	for _, str := range strs {
		charArray := []rune(str)
		sort.Slice(charArray, func(i, j int) bool {
			return charArray[i] < charArray[j]
		})
		var sortStr = string(charArray)
		_, ok := resMap[sortStr]

		if ok {
			resMap[sortStr] = append(resMap[sortStr], str)
		} else {
			resMap[sortStr] = []string{str}
		}
	}

	for key := range resMap {
		res = append(res, resMap[key])
	}

	return res
}

func (p P49)sol2(strs []string) [][]string {
	return nil
}

func main()  {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}

