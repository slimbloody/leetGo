package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type P145 struct {
}

func postorderTraversal(root *TreeNode) []int {
    var p P145
    return p.sol1(root)
}

func (p P145)sol1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// ...: a variable number of arguments
	res := make([]int, 0)
	res = append(res, p.sol1(root.Left)...)
	res = append(res, p.sol1(root.Right)...)
	res = append(res, root.Val)
	return res
}

func (p P145)sol2(root *TreeNode) []int {
}
