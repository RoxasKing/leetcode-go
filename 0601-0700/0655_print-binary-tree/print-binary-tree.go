package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if height < depth {
			height = depth
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	m, n := height+1, (1<<(height+1))-1
	o := make([][]string, m)
	for i := range o {
		o[i] = make([]string, n)
	}
	var fill func(node *TreeNode, x, y int)
	fill = func(node *TreeNode, x, y int) {
		o[x][y] = strconv.Itoa(node.Val)
		if node.Left != nil {
			fill(node.Left, x+1, y-(1<<(height-x-1)))
		}
		if node.Right != nil {
			fill(node.Right, x+1, y+(1<<(height-x-1)))
		}
	}
	fill(root, 0, (n-1)/2)
	return o
}
