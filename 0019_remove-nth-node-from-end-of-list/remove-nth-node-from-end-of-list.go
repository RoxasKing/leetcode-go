package main

/*
  给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

  说明：
   给定的 n 保证是有效的。

  进阶：
    你能尝试使用一趟扫描实现吗？
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeA, nodeB := head, head
	preA := &ListNode{Next: head}
	preH := preA
	for i := 0; i < n && nodeB != nil; i++ {
		nodeB = nodeB.Next
	}
	for nodeB != nil {
		nodeA = nodeA.Next
		nodeB = nodeB.Next
		preA = preA.Next
	}
	preA.Next = nodeA.Next
	return preH.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
