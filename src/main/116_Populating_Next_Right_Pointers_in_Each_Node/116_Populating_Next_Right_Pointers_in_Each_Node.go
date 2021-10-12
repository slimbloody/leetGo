package _16_Populating_Next_Right_Pointers_in_Each_Node

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

type P116 struct {
}

func connect(root *Node) *Node {
	var p P116
	return p.sol2(root)
}

func (p P116)sol1(root *Node) *Node {
	if root == nil {
		return nil
	}

	res := root
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		newQueue := make([]*Node, 0)

		pre := &queue[0]
		(*pre).Next = nil

		for i, node := range queue {
			if i != 0 {
				(*pre).Next = node
				pre = &(*pre).Next
				(*pre).Next = nil
			}

			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}

			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}

	return res
}


func (p P116)sol2(root *Node) *Node {
	if root == nil {
		return root
	}

	leftMost := root

	for leftMost != nil {
		head := leftMost

		for head != nil && head.Left != nil {
			head.Left.Next = head.Right

			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}

		leftMost = leftMost.Left
	}

	return root
}




