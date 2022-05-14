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
		return nil
	}
	for q := []*Node{root}; len(q) > 0; {
		n := len(q)
		for i := 0; i < n; i++ {
			if i+1 < n {
				q[i].Next = q[i+1]
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
	}
	return root
}
