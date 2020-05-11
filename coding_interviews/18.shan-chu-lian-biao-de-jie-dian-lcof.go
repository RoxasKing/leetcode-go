package codinginterviews

/*
  给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
  返回删除后的链表的头节点。
  注意：此题对比原题有改动

  说明：
  题目保证链表中节点的值互不相同
  若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
*/

func deleteNode(head *ListNode, val int) *ListNode {
	pre := &ListNode{Val: -1, Next: head}
	for node := pre; node != nil; node = node.Next {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			break
		}
	}
	return pre.Next
}
