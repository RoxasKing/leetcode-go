package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(head *ListNode) {
	var rev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = rev
		rev = cur
		cur = tmp
	}
	head = rev
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
