package main

/*
  请判断一个链表是否为回文链表。

  进阶：
    你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

// Double Pointer
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	slowPre := &ListNode{Next: head}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		slowPre = slowPre.Next
	}
	slowPre.Next = nil
	var reve *ListNode
	if fast == nil {
		reve = reverse(slow)
	} else {
		reve = reverse(slow.Next)
	}
	res := checkIfEqual(head, reve)
	if fast == nil {
		slowPre.Next = reverse(reve)
	} else {
		slowPre.Next = slow
		slow.Next = reverse(reve)
	}
	return res
}

func reverse(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

func checkIfEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	if l1 != nil || l2 != nil {
		return false
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
