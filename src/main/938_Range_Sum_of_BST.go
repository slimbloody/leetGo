package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type P938 struct {
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var p P938
	return p.sol1optimize1(root, low, high)
}

func (p P938)inorderTraversal(root *TreeNode, inorderTraversalArr *[]int) {
	if root == nil {
		return
	}
	p.inorderTraversal(root.Left, inorderTraversalArr)
	*inorderTraversalArr = append(*inorderTraversalArr, root.Val)
	p.inorderTraversal(root.Right, inorderTraversalArr)
}

func (p P938) sol1optimize1(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var res = 0

	if root.Val < high {
		res += p.sol1optimize1(root.Right, low, high)
	}

	if root.Val > low {
		res += p.sol1optimize1(root.Left, low, high)
	}

	if root.Val >= low && root.Val <= high {
		res += root.Val
	}

	return res
}

func (p P938) sol1optimize2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return p.sol1optimize2(root.Left, low, high)
	}

	if root.Val < low {
		return p.sol1optimize2(root.Right, low, high)
	}

	return root.Val + p.sol1optimize2(root.Left, low, high) + p.sol1optimize2(root.Right, low, high)
}

func (p P938)sol1(root *TreeNode, low int, high int) int {
	var inorderTraversalArr []int
	p.inorderTraversal(root, &inorderTraversalArr)
	var res = 0

	for _, i := range inorderTraversalArr {
		if i >= low {
			res += i
		}
		if i >= high {
			break
		}
	}

	return res
}


