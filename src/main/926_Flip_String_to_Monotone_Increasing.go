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

type P926 struct {
}

func minFlipsMonoIncr(s string) int {
	var p P926
	return p.sol1(s)
}

func (p P926)sol1(s string) int {
	var dp = make([]int, len(s) + 1)
	var res = int(^uint(0) >> 1)

	for i := 0; i < len(dp); i++ {
		var temp = i - dp[i] + (len(dp) - i - (dp[len(dp) - 1] - dp[i]))
		//var temp = len(dp) - dp[len(dp) - 1]
		if temp < res {
			res = temp
		}
	}

	return res
}

func main()  {
	// 1
	// [0 0 0 1 2 2]
	fmt.Println(minFlipsMonoIncr("00110"))
	// 2
	// [0 0 1 1 2 3 3]
	fmt.Println(minFlipsMonoIncr("010110"))
	// 2
	// [0 0 0 0 1 2 2 2 2]
	fmt.Println(minFlipsMonoIncr("00011000"))
}
