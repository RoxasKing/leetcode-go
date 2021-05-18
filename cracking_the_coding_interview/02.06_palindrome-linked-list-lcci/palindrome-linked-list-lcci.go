package main

/*
  Implement a function to check if a linked list is a palindrome.

  Example 1:
    Input:  1->2
    Output:  false

  Example 2:
    Input:  1->2->2->1
    Output:  true

  Follow up:
    Could you do it in O(n) time and O(1) space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-linked-list-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l1, l2 := head, reverse(slow.Next)
	slow.Next = nil
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var out *ListNode
	for head != nil {
		next := head.Next
		head.Next = out
		out = head
		head = next
	}
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}
