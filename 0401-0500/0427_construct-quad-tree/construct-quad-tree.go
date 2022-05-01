package main

// Difficulty:
// Medium

// Tags:
// Divide and Conquer
// DFS

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var dfs func(int, int, int) *Node
	dfs = func(x, y, b int) *Node {
		node := &Node{Val: true, IsLeaf: false}
		if b == 1 {
			node.Val = grid[x][y] == 1
			node.IsLeaf = true
			return node
		}
		b >>= 1
		tl := dfs(x, y, b)
		tr := dfs(x, y+b, b)
		bl := dfs(x+b, y, b)
		br := dfs(x+b, y+b, b)
		if tl.Val == tr.Val && tl.Val == bl.Val && tl.Val == br.Val {
			node.Val = tl.Val
			node.IsLeaf = tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf
		}
		if !node.IsLeaf {
			node.TopLeft = tl
			node.TopRight = tr
			node.BottomLeft = bl
			node.BottomRight = br
		}
		return node
	}
	return dfs(0, 0, len(grid))
}
