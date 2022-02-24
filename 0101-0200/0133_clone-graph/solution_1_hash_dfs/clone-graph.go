package main

// Difficulty:
// Medium

// Tags:
// Hash
// DFS

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	h := map[int]*Node{}
	vis := [101]bool{}
	var dfs func(*Node)
	dfs = func(node *Node) {
		if vis[node.Val] {
			return
		}
		vis[node.Val] = true
		for _, next := range node.Neighbors {
			if _, ok := h[next.Val]; !ok {
				h[next.Val] = &Node{Val: next.Val}
			}
			h[node.Val].Neighbors = append(h[node.Val].Neighbors, h[next.Val])
			dfs(next)
		}
	}
	h[node.Val] = &Node{Val: node.Val}
	dfs(node)
	return h[node.Val]
}
