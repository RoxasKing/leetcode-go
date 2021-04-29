package main

/*
  Given the root of a binary tree, determine if it is a complete binary tree.

  In a complete binary tree, every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

  Example 1:
    Input: root = [1,2,3,4,5,6]
    Output: true
    Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.

  Example 2:
    Input: root = [1,2,3,4,5,null,7]
    Output: false
    Explanation: The node with value 7 isn't as far left as possible.

  Constraints:
    1. The number of nodes in the tree is in the range [1, 100].
    2. 1 <= Node.val <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	flg := false
	for len(q) > 0 {
		size := len(q)
		for _, t := range q {
			if t == nil {
				flg = true
				continue
			}
			if flg {
				return false
			}
			q = append(q, t.Left, t.Right)
		}
		q = q[size:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
