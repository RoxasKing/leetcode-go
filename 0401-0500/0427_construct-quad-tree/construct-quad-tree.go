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
	var dfs func(x, y, d int) *Node
	dfs = func(x, y, d int) *Node {
		if d == 1 {
			return &Node{Val: grid[x][y] == 1, IsLeaf: true}
		}
		d /= 2
		tl := dfs(x, y, d)
		tr := dfs(x, y+d, d)
		bl := dfs(x+d, y, d)
		br := dfs(x+d, y+d, d)
		node := &Node{Val: tl.Val || tr.Val || bl.Val || br.Val}
		if tl.Val == tr.Val && tl.Val == bl.Val && bl.Val == br.Val {
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
