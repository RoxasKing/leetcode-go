package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 将链表存入栈中
	var stack []*ListNode
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	// 获取链表总长度
	n := len(stack)
	// 建立链表头指针
	head = &ListNode{Val: -1, Next: nil}
	// 初始化尾节点
	tail := head
	// 计算剩余不够一组的数量
	j := n % k
	// 将这些节点正常顺序插入到头节点之后
	for i := 0; i < j; i++ {
		// 获取栈顶节点
		node := stack[n-1]
		// 新节点连接头指针的后继节点
		node.Next = head.Next
		// 头指针连接新节点
		head.Next = node
		// 尾指针指向新节点
		tail = node
		n--
	}
	// 翻转节点流程
	for n > 0 {
		// 获取栈顶节点
		node := stack[n-1]
		// 新节点连接头指针的后继节点
		node.Next = head.Next
		// 头指针连接新节点
		head.Next = node
		// 尾指针指向新节点
		tail = node
		n--
		// 翻转当前组节点
		for i := 0; i < k-1; i++ {
			// 获取栈顶节点
			node := stack[n-1]
			// 新节点连接尾指针的后继节点
			node.Next = tail.Next
			// 尾指针连接新节点
			tail.Next = node
			// 尾指针指向新节点
			tail = node
			n--
		}
	}
	return head.Next
}

func main() {
	head1 := &ListNode{Val: 1}
	nex1 := &ListNode{Val: 2}
	head1.Next = nex1
	nex2 := &ListNode{Val: 3}
	nex1.Next = nex2
	nex3 := &ListNode{Val: 4}
	nex2.Next = nex3
	nex4 := &ListNode{Val: 5}
	nex3.Next = nex4

	new := reverseKGroup(head1, 2)

	for node := new; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}

}
