package main

import (
	"fmt"
	"strings"
)

type P1678 struct {
}

func interpret(command string) string {
	var p P1678
	return p.sol1(command)
}

func (p P1678)sol1(command string) string {
	return strings.Replace(
		strings.Replace(command, "()", "o", -1),
		"(al)",
		"al",
		-1,
		)
}

func (p P1678)sol2(command string) string {
	ans := ""
	for i := 0; i < len(command); {
		if command[i] == 'G'{
			ans += "G"
			i++
		} else if command[i + 1] == ')' {
			ans += "o"
			i += 2
		} else {
			ans += "al"
			i += 4
		}
	}
	return ans
}

func main() {
	fmt.Println(interpret(""))
	fmt.Println(interpret("G()(al)"))
	fmt.Println(interpret("G()()()()(al)"))
	fmt.Println(interpret("(al)G(al)()()G"))
}