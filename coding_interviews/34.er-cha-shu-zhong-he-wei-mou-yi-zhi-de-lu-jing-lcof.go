package codinginterviews

/*
  输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

  提示：
    节点总数 <= 10000
*/

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var (
		out    [][]int
		cur    []int
		curSum int
		dfs    func(*TreeNode)
	)
	dfs = func(root *TreeNode) {
		cur = append(cur, root.Val)
		curSum += root.Val
		if root.Left == nil && root.Right == nil && curSum == sum {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
		curSum -= root.Val
		cur = cur[:len(cur)-1]
	}
	dfs(root)
	return out
}
