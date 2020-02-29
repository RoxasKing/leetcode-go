package leetcode

/*
  给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
  例如，从根到叶子节点路径 1->2->3 代表数字 123。
  计算从根到叶子节点生成的所有数字之和。
  说明: 叶子节点是指没有子节点的节点。
*/

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	cur := root.Val
	var recur func(*TreeNode)
	recur = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			sum += cur
			return
		}
		if root.Left != nil {
			cur = cur*10 + root.Left.Val
			recur(root.Left)
			cur = (cur - root.Left.Val) / 10
		}
		if root.Right != nil {
			cur = cur*10 + root.Right.Val
			recur(root.Right)
			cur = (cur - root.Right.Val) / 10
		}
	}
	recur(root)
	return sum
}
