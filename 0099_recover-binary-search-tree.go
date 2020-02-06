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
		cur    = root
		pre    *TreeNode
		first  *TreeNode
		second *TreeNode
	)
	for cur != nil {
		if cur.Left != nil {
			curPre := cur.Left
			for curPre.Right != nil && curPre.Right != cur {
				curPre = curPre.Right
			}
			if curPre.Right != cur {
				curPre.Right = cur
				cur = cur.Left
			} else {
				curPre.Right = nil
				if pre != nil && pre.Val > cur.Val {
					if first == nil {
						first = pre
					}
					second = cur
				}
				pre = cur
				cur = cur.Right
			}
		} else {
			if pre != nil && pre.Val > cur.Val {
				if first == nil {
					first = pre
				}
				second = cur
			}
			pre = cur
			cur = cur.Right
		}
	}
	first.Val, second.Val = second.Val, first.Val
}