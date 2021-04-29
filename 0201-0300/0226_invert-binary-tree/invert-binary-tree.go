package main

/*
  Given the root of a binary tree, invert the tree, and return its root.

  Example 1:
    Input: root = [4,2,7,1,3,6,9]
    Output: [4,7,2,9,6,3,1]

  Example 2:
    Input: root = [2,1,3]
    Output: [2,3,1]

  Example 3:
    Input: root = []
    Output: []

  Constraints:
    1. The number of nodes in the tree is in the range [0, 100].
    2. -100 <= Node.val <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/invert-binary-tree
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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
