package main

import "fmt"

type P1662 struct {
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var p P1662
	return p.sol1(word1, word2)
}

func (p P1662)sol1(word1 []string, word2 []string) bool {
	channel1 := p.sol1GetChannel(word1)
	channel2 := p.sol1GetChannel(word2)

	if len(channel1) != len(channel2) {
		return false
	}

	var length = len(channel1)
	for i := 0; i < length; i++ {
		c1, _ := <- channel1
		c2, _ := <- channel2
		if c1 != c2 {
			return false
		}
	}
	//wxgdwznoledlfeoilewsjziotnncyebhwpdnnimcorzovuiiglycfb
	//wxgdwznoledlfeoilewsjziotnncyebhwpdnnimcoriiglycfbhnjm

	return true
}

func (p P1662)sol1GetChannel(strs []string) chan rune {
	c := make(chan rune, 1000000)
	defer close(c)
	for _, str := range strs {
		runes := []rune(str)
		for _, runeElem := range runes {
			c <- runeElem
		}
	}
	return c
}

// sol2: double pointer, pointer1 for string, pointer2 for rune
func (p P1662)sol2(word1 []string, word2 []string) bool {
	return true
}

func main()  {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
	fmt.Println(arrayStringsAreEqual(
		[]string{"wxgdwznoledlfeoilewsjziotnncyebhwpdnnimcorzovuiig","lycfb"},
		[]string{"wxgdwznoledlfeoilewsjzio","tnncyebhwpdnnimcor","iigl","yc","f","b","hnjm"}))
}
