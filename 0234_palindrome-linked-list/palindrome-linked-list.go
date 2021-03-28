package main

/*
  Given the head of a singly linked list, return true if it is a palindrome.

  Example 1:
    Input: head = [1,2,2,1]
    Output: true

  Example 2:
    Input: head = [1,2]
    Output: false

  Constraints:
    1. The number of nodes in the list is in the range [1, 10^5].
    2. 0 <= Node.val <= 9

  Follow up: Could you do it in O(n) time and O(1) space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-linked-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
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
