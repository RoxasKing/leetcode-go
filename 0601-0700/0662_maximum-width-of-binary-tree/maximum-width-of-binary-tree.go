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
	out, cur_d, l_p := 0, 0, 0
	for q := []*pair{{root, 0, 0}}; len(q) > 0; q = q[1:] {
		e := q[0]
		node, d, p := e.node, e.d, e.p
		if node == nil {
			continue
		}
		if cur_d != d {
			cur_d, l_p = d, p
		}
		out = Max(out, p+1-l_p)
		q = append(q, &pair{node.Left, d + 1, p << 1}, &pair{node.Right, d + 1, p<<1 + 1})
	}
	return out
}

type pair struct {
	node *TreeNode
	d, p int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
