package main

import "fmt"

func ReverseString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s) / 2; i++ {
		runes[i], runes[len(s) - 1 - i] = runes[len(s) - 1 - i], runes[i]
	}

	return string(runes)
}

func processLeftStr(index *int, num string, carry *int, res string) string {
	for (*index) < len(num) {
		var temp = int(num[(*index)] - '0') + (*carry)
		*index++
		if temp > 9 {
			*carry = temp / 10
			temp %= 10
		} else {
			*carry = 0
		}
		res += string(rune(temp + int('0')))

	}
	return res
}


func addStrings(num1 string, num2 string) string {
	var res string = ""
	len1, len2 := len(num1), len(num2)
	num1 = ReverseString(num1)
	num2 = ReverseString(num2)
	carry := 0
	index := 0

	for index < len1 && index < len2 {
		var temp int = int(num1[index]) - int('0') + int(num2[index]) - int('0') + carry
		index++
		if temp > 9 {
			carry = temp / 10
			temp %= 10
		} else {
			carry = 0
		}
		res += string(rune(temp + int('0')))
	}

	res = processLeftStr(&index, num1, &carry, res)
	res = processLeftStr(&index, num2, &carry, res)

	if carry > 0 {
		res += string('1')
	}

	res = ReverseString(res)
	return res
}

func main() {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
	fmt.Println(addStrings("1", "9"))
	fmt.Println(addStrings("999", "999"))
}