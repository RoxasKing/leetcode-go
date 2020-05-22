package codinginterviews

/*
  输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

  提示：
    节点总数 <= 10000
*/

// DFS + Backtrack
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var (
		res    [][]int
		cur    []int
		curSum int
		dfs    func(*TreeNode)
	)
	pushNode := func(node *TreeNode) {
		curSum += node.Val
		cur = append(cur, node.Val)
	}
	popNode := func(node *TreeNode) {
		curSum -= node.Val
		cur = cur[:len(cur)-1]
	}
	addToResult := func() {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)
	}
	dfs = func(node *TreeNode) {
		pushNode(node)
		if curSum == sum && isLeaf(node) {
			addToResult()
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		popNode(node)
	}
	dfs(root)
	return res
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
