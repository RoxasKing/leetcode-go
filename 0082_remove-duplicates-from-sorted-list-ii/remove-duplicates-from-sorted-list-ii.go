package main

/*
  Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

  Example 1:
    Input: head = [1,2,3,3,4,4,5]
    Output: [1,2,5]

  Example 2:
    Input: head = [1,1,1,2,3]
    Output: [2,3]

  Constraints:
    1. The number of nodes in the list is in the range [0, 300].
    2. -100 <= Node.val <= 100
    3. The list is guaranteed to be sorted in ascending order.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	ptr := headPre
	isDup := false
	for n := head.Next; n != nil; n = n.Next {
		if n.Val == ptr.Next.Val {
			isDup = true
		} else {
			if isDup {
				ptr.Next = n
			} else {
				ptr = ptr.Next
			}
			isDup = false
		}
	}
	if isDup {
		ptr.Next = nil
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
