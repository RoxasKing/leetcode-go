package main

/*
  Given the root of a binary tree, return the length of the diameter of the tree.

  The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

  The length of a path between two nodes is represented by the number of edges between them.

  Example 1:
    Input: root = [1,2,3,4,5]
    Output: 3
    Explanation: 3is the length of the path [4,2,1,3] or [5,2,1,3].

  Example 2:
    Input: root = [1,2]
    Output: 1

  Constraints:
    1. The number of nodes in the tree is in the range [1, 10^4].
    2. -100 <= Node.val <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
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
func diameterOfBinaryTree(root *TreeNode) int {
	out := 0
	dfs(root, 0, &out)
	return out
}

func dfs(root *TreeNode, depth int, out *int) int {
	if root == nil {
		return depth
	}

	depth++
	ld := dfs(root.Left, depth, out)
	rd := dfs(root.Right, depth, out)
	*out = Max(*out, ld+rd-depth<<1)
	return Max(ld, rd)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
