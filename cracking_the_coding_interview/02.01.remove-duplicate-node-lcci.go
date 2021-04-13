package main

/*
  编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

  提示：
    链表长度在[0, 20000]范围内。
    链表元素在[0, 20000]范围内。

  进阶：
    如果不得使用临时缓冲区，该怎么解决？
*/

// Hash
func removeDuplicateNodes(head *ListNode) *ListNode {
	mark := make(map[int]struct{})
	pre := &ListNode{Val: -1, Next: head}
	cur := pre
	for cur != nil && cur.Next != nil {
		if _, ok := mark[cur.Next.Val]; !ok {
			mark[cur.Next.Val] = struct{}{}
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}
	}
	return pre.Next
}

// Enumerate + Double Loop
func removeDuplicateNodes2(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		ptr := cur
		for ptr.Next != nil {
			if ptr.Next.Val == cur.Val {
				ptr.Next = ptr.Next.Next
			} else {
				ptr = ptr.Next
			}
		}
		cur = cur.Next
	}
	return head
}
