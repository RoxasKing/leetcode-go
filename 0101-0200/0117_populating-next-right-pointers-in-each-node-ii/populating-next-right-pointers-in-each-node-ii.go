package main

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{root}
	for len(q) != 0 {
		curSize := len(q)
		for i := 0; i < curSize; i++ {
			if i > 0 {
				q[i-1].Next = q[i]
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[curSize:]
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
