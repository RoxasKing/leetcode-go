package main

// Difficulty:
// Medium

// Tags:
// Hash
// BFS

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	dict := make(map[int]*Node)
	dict[node.Val] = &Node{Val: node.Val}
	q := []*Node{node}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, n := range node.Neighbors {
			if _, ok := dict[n.Val]; !ok {
				dict[n.Val] = &Node{Val: n.Val}
				q = append(q, n)
			}
			dict[node.Val].Neighbors = append(dict[node.Val].Neighbors, dict[n.Val])
		}
	}
	return dict[node.Val]
}
