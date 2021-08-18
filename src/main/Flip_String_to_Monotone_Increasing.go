package main

import "fmt"

/*
//Go不支持运算符重载，因此需要先将 a<b 在函数外转换为 bool 条件
//Go不支持泛型，只能用 interface{} 模拟
//返回的类型安全需要用户自己保证，.(type) 的类型必须匹配
//interface{} 是运行时泛型，性能没有编译时泛型高
func If(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}

	return b
}
 */

func minFlipsMonoIncr(s string) int {
	res := 0
	start, end := 0, len(s) - 1

	for start < end {
		if s[start] == '0' {
			start++
		} else {
			break
		}
	}

	//var dp [end - start +1]int
	for start < end {
		if s[end] == '1' {
			end--
		} else {
			break
		}
	}
	for i, c := start, 0; i <= end; i, c = i + 1, c + 1 {
		temp := 0
		for {
		}
	}

	if start < end {
	} else {
		return res
	}
}

func main()  {
	fmt.Println(minFlipsMonoIncr("00110"))
	fmt.Println(minFlipsMonoIncr("010110"))
	fmt.Println(minFlipsMonoIncr("00011000"))
}
