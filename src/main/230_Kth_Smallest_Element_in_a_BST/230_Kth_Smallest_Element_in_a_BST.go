package _30_Kth_Smallest_Element_in_a_BST

type P230 struct {
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func kthSmallest(root *TreeNode, k int) int {
    var p P230
    return p.sol1(root, k)
}

func (p P230)sol1(root *TreeNode, k int) int {
    count := k
    inorder(root, count)
}

func inorder(root *TreeNode, count int) int {
	inorder(root.Left, count)

    inorder(root.Right, count)
}


