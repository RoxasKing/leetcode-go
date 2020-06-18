package leetcode

import "strconv"

/*
  我们从二叉树的根节点 root 开始进行深度优先搜索。
  在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。
  如果节点只有一个子节点，那么保证该子节点为左子节点。
  给出遍历输出 S，还原树并返回其根节点 root。

  提示：
    原始树中的节点数介于 1 和 1000 之间。
    每个节点的值介于 1 和 10 ^ 9 之间。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursive
func recoverFromPreorder(S string) *TreeNode {
	var dfs func(int) *TreeNode
	dfs = func(depth int) *TreeNode {
		var l int
		for l < len(S) && S[l] == '-' {
			l++
		}
		if depth > l || l == len(S) {
			return nil
		}
		r := l
		for r < len(S) && S[r] != '-' {
			r++
		}
		rootVal, _ := strconv.Atoi(S[l:r])
		S = S[r:]
		lNode := dfs(depth + 1)
		rNode := dfs(depth + 1)
		return &TreeNode{
			Val:   rootVal,
			Left:  lNode,
			Right: rNode,
		}
	}
	return dfs(0)
}
