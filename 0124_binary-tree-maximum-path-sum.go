package leetcode

/*
  给定一个非空二叉树，返回其最大路径和。
  本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/

// Recursion
func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 31
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := Max(maxGain(node.Left), 0)
		rightGain := Max(maxGain(node.Right), 0)
		maxSum = Max(maxSum, node.Val+leftGain+rightGain)
		return node.Val + Max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
