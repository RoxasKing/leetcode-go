package My_LeetCode_In_Go

/*
  给定一个二叉树，找出其最大深度。
  二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
  说明: 叶子节点是指没有子节点的节点。
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
