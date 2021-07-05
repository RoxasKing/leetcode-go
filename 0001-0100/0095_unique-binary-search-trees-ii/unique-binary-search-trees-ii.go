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
func generateTrees(n int) []*TreeNode {
	return dfs(1, n)
}

func dfs(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	out := []*TreeNode{}
	for i := l; i <= r; i++ {
		ls := dfs(l, i-1)
		rs := dfs(i+1, r)
		for _, l := range ls {
			for _, r := range rs {
				out = append(out, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
