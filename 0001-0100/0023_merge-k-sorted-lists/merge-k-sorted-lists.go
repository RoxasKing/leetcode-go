package main

/*
  给你一个链表数组，每个链表都已经按升序排列。
  请你将所有链表合并到一个升序链表中，返回合并后的链表。

  示例 1：
    输入：lists = [[1,4,5],[1,3,4],[2,6]]
    输出：[1,1,2,3,4,4,5,6]
    解释：链表数组如下：
    [
      1->4->5,
      1->3->4,
      2->6
    ]
    将它们合并到一个有序链表中得到。
    1->1->2->3->4->4->5->6

  示例 2：
    输入：lists = []
    输出：[]

  示例 3：
    输入：lists = [[]]
    输出：[]

  提示：
    k == lists.length
    0 <= k <= 10^4
    0 <= lists[i].length <= 500
    -10^4 <= lists[i][j] <= 10^4
    lists[i] 按 升序 排列
    lists[i].length 的总和不超过 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Iteration
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	newLists := make([]*ListNode, len(lists))
	copy(newLists, lists)
	for base := 1; base <= len(lists); base *= 2 {
		for i := 0; i < len(newLists); i += 2 * base {
			if i+base < len(lists) {
				newLists[i] = mergeTwoLists(newLists[i], newLists[i+base])
			}
		}
	}
	return newLists[0]
}

// Recursion
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return recurMergeKLists(&lists, 0, len(lists)-1)
}

func recurMergeKLists(lists *[]*ListNode, l, r int) *ListNode {
	if l == r {
		return (*lists)[l]
	}
	m := l + (r-l)>>1
	return mergeTwoLists(recurMergeKLists(lists, l, m), recurMergeKLists(lists, m+1, r))
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	pre := new(ListNode)
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
