package main

/*
  给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

  说明：
  你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

  进阶：
    如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Morris Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			if pre.Right != root {
				pre.Right = root
				root = root.Left
				continue
			}
			pre.Right = nil
		}
		k--
		if k == 0 {
			break
		}
		root = root.Right
	}
	return root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
