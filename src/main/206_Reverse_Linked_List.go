package main


type P206 struct {
}

func reverseList(head *ListNode) *ListNode {
	var p P206
	return p.sol1(head)
}

func (p P206)sol1(head *ListNode) *ListNode {
	res := &ListNode{}
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = nil

		if res.Next == nil {
			res.Next = cur
		} else {
			cur.Next = res.Next
			res.Next = cur
		}
	}

	return res.Next
}

