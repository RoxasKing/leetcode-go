package codinginterviews

/*
  输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

  限制：
    1 <= 树的结点个数 <= 10000
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if Abs(l-r) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
