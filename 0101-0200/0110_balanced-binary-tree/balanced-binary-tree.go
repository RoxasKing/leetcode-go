package main

/*
  Given a binary tree, determine if it is height-balanced.

  For this problem, a height-balanced binary tree is defined as:

    a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

  Example 1:
    Input: root = [3,9,20,null,null,15,7]
    Output: true

  Example 2:
    Input: root = [1,2,2,3,3,null,null,4,4]
    Output: false

  Example 3:
    Input: root = []
    Output: true

  Constraints:
    1. The number of nodes in the tree is in the range [0, 5000].
    2. -10^4 <= Node.val <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/balanced-binary-tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && Abs(dfs(root.Left)-dfs(root.Right)) < 2
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(dfs(root.Left), dfs(root.Right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
