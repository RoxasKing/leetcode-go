package main

// Difficulty:
// Medium

// Tags:
// BFS

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
			if i > 0 {
				q[i-1].Next = q[i]
			}
		}
		q = q[n:]
	}
	return root
}
