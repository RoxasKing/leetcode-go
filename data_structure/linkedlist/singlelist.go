package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(head *ListNode) *ListNode {
	var rev *ListNode
	for cur := head; cur != nil; {
		tmp := cur.Next
		cur.Next, rev, cur = rev, cur, tmp
	}
	return rev
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
