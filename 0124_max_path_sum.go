package My_LeetCode_In_Go

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		max(maxPathSum(root.Left), maxPathSum(root.Right))+root.Val,
		root.Val,
	)
}
