package main

/*
  Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

  Note: This question is the same as 530: https://leetcode.com/problems/minimum-absolute-difference-in-bst/

  Example 1:
    Input: root = [4,2,6,1,3]
    Output: 1

  Example 2:
    Input: root = [1,0,48,null,null,12,49]
    Output: 1

  Constraints:
    1. The number of nodes in the tree is in the range [2, 100].
    2. 0 <= Node.val <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	out := 1<<31 - 1
	preVal := -1
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
		if preVal != -1 && root.Val-preVal < out {
			out = root.Val - preVal
		}
		preVal = root.Val
		root = root.Right
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
