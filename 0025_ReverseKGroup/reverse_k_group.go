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
	var out *ListNode
	pre := &ListNode{}
	pre.Next = head
	out = pre

	for {
		i := 0
		tmp1 := head
		for ; i < k-1; i++ {
			if tmp1 != nil {
				tmp1 = tmp1.Next
			} else {
				break
			}
		}
		if i < k-1 || tmp1 == nil {
			return out.Next
		} else {
			end := tmp1
			next := end.Next

			tmp2 := end
			for i := 0; i < k-2; i++ {
				node := head
				for node.Next != tmp2 {
					node = node.Next
				}
				tmp2.Next = node
				tmp2 = tmp2.Next
			}
			tmp2.Next = head
			head.Next = next
			pre.Next = end

			for i := 0; i < k; i++ {
				if pre != nil {
					pre = pre.Next
				} else {
					break
				}
			}
			head = pre.Next
		}
	}
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
