package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util"
	. "My_LeetCode_In_Go/util/tree"
)

/*
  给定一个非空二叉树，返回其最大路径和。
  本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/

func maxPathSum(root *TreeNode) int {
	max := -1 << 31
	maxGain(root, &max)
	return max
}

func maxGain(node *TreeNode, max_sum *int) int {
	if node == nil {
		return 0
	}
	left_gain := Max(maxGain(node.Left, max_sum), 0)
	right_gain := Max(maxGain(node.Right, max_sum), 0)
	*max_sum = Max(*max_sum, node.Val+left_gain+right_gain)
	return node.Val + Max(left_gain, right_gain)
}
