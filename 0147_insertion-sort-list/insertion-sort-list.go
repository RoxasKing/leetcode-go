package main

/*
  对链表进行插入排序。

  插入排序算法：
    插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
    每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
    重复直到所有输入数据插入完为止。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/insertion-sort-list
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	last := head
	node := head.Next
	head.Next = nil
	for node != nil {
		nodeNext := node.Next
		if node.Val > last.Val {
			last.Next = node
			last = last.Next
			last.Next = nil
			node = nodeNext
			continue
		}
		ptr := headPre
		for ptr.Next != nil && ptr.Next.Val <= node.Val {
			ptr = ptr.Next
		}
		ptrNext := ptr.Next
		ptr.Next = node
		node.Next = ptrNext
		node = nodeNext
	}
	return headPre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
