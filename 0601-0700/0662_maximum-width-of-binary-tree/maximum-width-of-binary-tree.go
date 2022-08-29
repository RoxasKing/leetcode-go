package main

// Difficulty:
// Medium

// Tags:
// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	o := 0
	for q := []*pair{{root, 1}}; len(q) > 0; {
		n := len(q)
		o = max(o, q[n-1].i+1-q[0].i)
		for _, p := range q {
			if p.t.Left != nil {
				q = append(q, &pair{p.t.Left, p.i << 1})
			}
			if p.t.Right != nil {
				q = append(q, &pair{p.t.Right, p.i<<1 + 1})
			}
		}
		q = q[n:]
	}
	return o
}

type pair struct {
	t *TreeNode
	i int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
