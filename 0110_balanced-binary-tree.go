package main

/*
  给定一个二叉树，判断它是否是高度平衡的二叉树。
  本题中，一棵高度平衡二叉树定义为：
    一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
  示例 1:
  给定二叉树 [3,9,20,null,null,15,7]
        3
       / \
      9  20
        /  \
       15   7
  返回 true 。
  示例 2:
  给定二叉树 [1,2,2,3,3,null,null,4,4]
           1
          / \
         2   2
        / \
       3   3
      / \
     4   4
  返回 false 。
*/

// Recursion : top down
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(*TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + Max(height(root.Left), height(root.Right))
	}
	return isBalanced(root.Left) && isBalanced(root.Right) &&
		Abs(height(root.Left)-height(root.Right)) < 2
}

// Recursion : bottom up
func isBalanced2(root *TreeNode) bool {
	var helper func(*TreeNode) (int, bool)
	helper = func(root *TreeNode) (height int, isBalanced bool) {
		if root == nil {
			return 0, true
		}
		leftHeight, leftBalance := helper(root.Left)
		if !leftBalance {
			return 0, false
		}
		rightHeight, rightBalance := helper(root.Right)
		if !rightBalance {
			return 0, false
		}
		if Abs(leftHeight-rightHeight) >= 2 {
			return 0, false
		}
		return Max(leftHeight, rightHeight) + 1, true
	}
	_, isBalanced := helper(root)
	return isBalanced
}
