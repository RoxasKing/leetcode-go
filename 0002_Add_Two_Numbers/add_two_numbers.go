package main

<<<<<<< HEAD
import "fmt"
=======
import (
	"fmt"
)
>>>>>>> fbd37d4face5d0a6b5a9061c15e8a182736e8a62

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 头节点
	resPre := &ListNode{}
	// 轮询节点
	cur := resPre
	// 进位数
	carry := 0

	// 只要链表其一有数，或者进位数不为 0 ，循环继续
	for l1 != nil || l2 != nil || carry > 0 {
		// 初始化下一节点起始数
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return resPre.Next
}

func main() {
	head1 := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	head1.Next = node1
	node2 := &ListNode{Val: 3}
	head1.Next = node2
	head2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	head2.Next = node3
	node4 := &ListNode{Val: 5}
	head1.Next = node4
	fmt.Println(addTwoNumbers(head1, head2))
}
