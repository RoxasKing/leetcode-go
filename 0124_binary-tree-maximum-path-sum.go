package My_LeetCode_In_Go

/*
  给定一个非空二叉树，返回其最大路径和。
  本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/

func maxPathSum(root *TreeNode) int {
	max := -1 << 31
	max_gain(root, &max)
	return max
}

func max_gain(node *TreeNode, max_sum *int) int {
	if node == nil {
		return 0
	}
	left_gain := max(max_gain(node.Left, max_sum), 0)
	right_gain := max(max_gain(node.Right, max_sum), 0)
	*max_sum = max(*max_sum, node.Val+left_gain+right_gain)
	return node.Val + max(left_gain, right_gain)
}
