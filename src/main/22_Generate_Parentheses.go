package main

type P22 struct {
}

func generateParenthesis(n int) []string {
	var p P22
	return p.sol1(n)
}

func (p P22)sol1(n int) []string {
	var res []string
	return p.sol1GenP(res, "", n, n)
}

func (p P22) sol1GenP(res []string, s string, left int, right int) []string {
	if left == 0 && right == 0 {
		res = append(res, s)
	}

	if left > 0 {
		if left <= right {
			p.sol1GenP(res, s + "(", left - 1, right)
			p.sol1GenP(res, s + ")", left, right - 1)
		} else {
		}
	} else {
	}
}

