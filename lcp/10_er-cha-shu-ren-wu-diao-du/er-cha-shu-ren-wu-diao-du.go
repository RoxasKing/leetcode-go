package main

// Tags:
// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimalExecTime(root *TreeNode) float64 {
	total, parallel := dfs(root)
	return total - parallel
}

func dfs(node *TreeNode) (float64, float64) {
	if node == nil {
		return 0, 0
	}

	total1, parallel1 := dfs(node.Left)
	total2, parallel2 := dfs(node.Right)

	if total1 < total2 {
		total1, total2 = total2, total1
		parallel1 = parallel2
	}

	if total1-2*parallel1 > total2 {
		return float64(node.Val) + total1 + total2, parallel1 + total2
	}
	return float64(node.Val) + total1 + total2, (total1 + total2) / 2 // (total1-total2)/2 + total2 = (total1+total2)/2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
