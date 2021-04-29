package main

/*
  A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

  The path sum of a path is the sum of the node's values in the path.

  Given the root of a binary tree, return the maximum path sum of any path.

  Example 1:
    Input: root = [1,2,3]
    Output: 6
    Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

  Example 2:
    Input: root = [-10,9,20,null,null,15,7]
    Output: 42
    Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

  Constraints:
    1. The number of nodes in the tree is in the range [1, 3 * 10^4].
    2. -1000 <= Node.val <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	out := -1 << 31
	dfs(root, &out)
	return out
}

func dfs(root *TreeNode, out *int) int {
	if root == nil {
		return 0
	}
	lSum := Max(dfs(root.Left, out), 0)
	rSum := Max(dfs(root.Right, out), 0)
	*out = Max(*out, root.Val+lSum+rSum)
	return root.Val + Max(lSum, rSum)
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
