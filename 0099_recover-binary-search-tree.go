package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util/tree"
)

/*
  二叉搜索树中的两个节点被错误地交换。
  请在不改变其结构的情况下，恢复这棵树。
*/

func recoverTree(root *TreeNode) {
	var (
		cur = root
		pre *TreeNode
		one *TreeNode
		two *TreeNode
	)
	for cur != nil {
		if cur.Left == nil {
			if pre != nil && pre.Val > cur.Val {
				if one == nil {
					one = pre
				}
				two = cur
			}
			pre = cur
			cur = cur.Right
		} else {
			leftRight := cur.Left
			for leftRight.Right != nil && leftRight.Right != cur {
				leftRight = leftRight.Right
			}
			if leftRight.Right != cur {
				leftRight.Right = cur
				cur = cur.Left
			} else {
				leftRight.Right = nil
				if pre != nil && pre.Val > cur.Val {
					if one == nil {
						one = pre
					}
					two = cur
				}
				pre = cur
				cur = cur.Right
			}
		}
	}
	one.Val, two.Val = two.Val, one.Val
}
