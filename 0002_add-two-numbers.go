package leetcode

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	var val, rest int
	for {
		val = rest
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val == 0 && l1 == nil && l2 == nil {
			if head.Next == nil {
				return head
			}
			return head.Next
		}
		if val >= 10 {
			val -= 10
			rest = 1
		} else {
			rest = 0
		}
		tail.Next = &ListNode{Val: val, Next: nil}
		tail = tail.Next
	}
}
