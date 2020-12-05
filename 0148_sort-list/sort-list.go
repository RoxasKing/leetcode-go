package main

/*
  给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

  进阶：
    你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

  示例 1：
    输入：head = [4,2,1,3]
    输出：[1,2,3,4]

  示例 2：
    输入：head = [-1,5,3,4,0]
    输出：[-1,0,3,4,5]

  示例 3：
    输入：head = []
    输出：[]

  提示：
    链表中节点的数目在范围 [0, 5 * 104] 内
    -105 <= Node.val <= 105

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sort-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Merge Sort + Linked List
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil
	return merge(sortList(l1), sortList(l2))
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
	}
	if l2 != nil {
		ptr.Next = l2
	}
	return outPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
