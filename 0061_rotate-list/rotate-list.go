package main

/*
  Given the head of a linked list, rotate the list to the right by k places.

  Example 1:
    Input: head = [1,2,3,4,5], k = 2
    Output: [4,5,1,2,3]

  Example 2:
    Input: head = [0,1,2], k = 4
    Output: [2,0,1]

  Constraints:
    1. The number of nodes in the list is in the range [0, 500].
    2. -100 <= Node.val <= 100
    3. 0 <= k <= 2 * 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rotate-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	headPre := &ListNode{Next: head}
	last := headPre
	size := 0
	for last.Next != nil {
		last = last.Next
		size++
	}
	if size == k {
		return head
	}
	last.Next = head // ring
	outPre := headPre
	for k := size - k%size; k > 0; k-- {
		outPre = outPre.Next
	}
	out := outPre.Next
	outPre.Next = nil
	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}
