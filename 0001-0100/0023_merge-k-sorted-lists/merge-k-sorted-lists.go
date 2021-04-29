package main

/*
  You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

  Merge all the linked-lists into one sorted linked-list and return it.

  Example 1:
    Input: lists = [[1,4,5],[1,3,4],[2,6]]
    Output: [1,1,2,3,4,4,5,6]
    Explanation: The linked-lists are:
      [
        1->4->5,
        1->3->4,
        2->6
      ]
      merging them into one sorted list:
      1->1->2->3->4->4->5->6

  Example 2:
    Input: lists = []
    Output: []

  Example 3:
    Input: lists = [[]]
    Output: []

  Constraints:
    1. k == lists.length
    2. 0 <= k <= 10^4
    3. 0 <= lists[i].length <= 500
    4. -10^4 <= lists[i][j] <= 10^4
    5. lists[i] is sorted in ascending order.
    6. The sum of lists[i].length won't exceed 10^4.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Merge Sort

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}
	m := n >> 1
	return merge(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
}

func merge(l1, l2 *ListNode) *ListNode {
	outPre := &ListNode{}
	ptr := outPre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}
	if l1 != nil {
		ptr.Next = l1
	} else if l2 != nil {
		ptr.Next = l2
	}
	return outPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
