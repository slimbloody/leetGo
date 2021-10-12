package _17_Populating_Next_Right_Pointers_in_Each_Node_II

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

type P117 struct {
}

func connect(root *Node) *Node {
	var p P117
	return p.connect(root)
}

func (p P117)connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	res := root
	queue := make([]*Node, 0)
	queue = append(queue, root)

	//jsonQ, _ := json.Marshal(queue)
	//fmt.Println(string(jsonQ))

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





