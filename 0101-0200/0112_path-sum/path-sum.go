package main

/*
  Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

  A leaf is a node with no children.

  Example 1:
    Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
    Output: true

  Example 2:
    Input: root = [1,2,3], targetSum = 5
    Output: false

  Example 3:
    Input: root = [1,2], targetSum = 0
    Output: false

  Constraints:
    The number of nodes in the tree is in the range [0, 5000].
    -1000 <= Node.val <= 1000
    -1000 <= targetSum <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/path-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	if root.Left != nil && hasPathSum(root.Left, targetSum) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, targetSum) {
		return true
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
