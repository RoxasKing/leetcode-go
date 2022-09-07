package main

// Difficulty:
// Medium

// Tags:
// BFS

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	o := [][]int{}
	for q := []*Node{root}; len(q) > 0; {
		a := []int{}
		n := len(q)
		for _, t := range q {
			a = append(a, t.Val)
			q = append(q, t.Children...)
		}
		q = q[n:]
		o = append(o, a)
	}
	return o
}
