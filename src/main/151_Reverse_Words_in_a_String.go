package main

import (
	"fmt"
	"strings"
)

type P151 struct {
}

func reverseWords(s string) string {
	var p P151
	return p.sol2(s)
}

func (p P151)sol1(s string) string {
	f := func(c rune) bool {
		return c == ' '
	}

	//fields := strings.Fields(s)
	fields := strings.FieldsFunc(s, f)
	//fmt.Println(reflect.TypeOf(fields))
	for i, j := 0, len(fields) - 1; i < j; i, j = i + 1, j - 1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	res := strings.Join(fields, " ")

	return res
}

// no trim, no split, no stringbuilder
func (p P151)sol2(s string) string {
	s = p.cleanSpace(s)
	s = p.reverseString(s)
	s = p.reverseWords(s)

	return s
}

func (p P151) reverseRuneWord(runes []rune, start int, end int) []rune {
	for ; start < end; start, end = start + 1, end - 1 {
		runes[start], runes[end] = runes[end], runes[start]
	}

	return runes
}

func (p P151)reverseWords(s string) string {
	runes := []rune(s)
	start := 0
	length := len(s)

	for start < length {
		j := start

		for j < length {
			if j == length - 1 {
				p.reverseRuneWord(runes, start, j)
				start = j + 1
				break
			}

			if s[j] != ' ' {
				j++
			} else {
				p.reverseRuneWord(runes, start, j - 1)
				start = j + 1
				break
			}
		}
	}

	return string(runes)
}

func (p P151) reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func (p P151)cleanSpace(s string) string {
	s = strings.Trim(s, " ")
	runeS := []rune(s)
	fullLen := len(runeS)
	i, j := 0, 0
	res := make([]rune, fullLen)

	for j < fullLen {
		if runeS[j] == ' ' && res[i] == ' ' {
			j++
		} else {
			res[i] = runeS[j]
			i++
			j++
		}
	}

	return string(res)
}

func main()  {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords("  Bob    Loves  Alice   "))
	fmt.Println(reverseWords("Alice does not even like bob"))
	fmt.Println(reverseWords("                            "))
}