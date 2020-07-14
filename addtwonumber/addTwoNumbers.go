package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	init := &ListNode{}
	head := init

	r := 0
	c1 := l1
	c2 := l2

	for c1 != nil || c2 != nil || r > 0 {
		sum := r
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}

		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}

		currentNode := &ListNode{
			Val: sum % 10,
		}

		r = sum / 10

		head.Next = currentNode
		head = currentNode
	}

	return init.Next
}
