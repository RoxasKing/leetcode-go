package My_LeetCode_In_Go

/*
  给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var size int
	for tmp := head; tmp != nil; tmp = tmp.Next {
		size++
	}
	index := size - n
	tmp := &ListNode{-1, nil}
	tmp.Next = head
	node := tmp
	for i := 0; i < index; i++ {
		node = node.Next
	}
	if node == nil || node.Next == nil {
		return nil
	} else if node.Next.Next == nil {
		node.Next = nil
	} else {
		node.Next = node.Next.Next
	}
	return tmp.Next
}
