package main

/*
  Given the roots of two binary trees p and q, write a function to check if they are the same or not.

  Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

  Example 1:
    Input: p = [1,2,3], q = [1,2,3]
    Output: true

  Example 2:
    Input: p = [1,2], q = [1,null,2]
    Output: false

  Example 3:
    Input: p = [1,2,1], q = [1,1,2]
    Output: false

  Constraints:
    1. The number of nodes in both trees is in the range [0, 100].
    2. -10^4 <= Node.val <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/same-tree
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

// DFS + Recursion
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
