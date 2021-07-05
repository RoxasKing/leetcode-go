package main

// Tags:
// BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	q := []*AnnotatedNode{{node: root, depth: 0, pos: 0}}
	out, curDepth, left := 0, 0, 0
	for len(q) > 0 {
		an := q[0]
		q = q[1:]
		node, depth, pos := an.node, an.depth, an.pos
		if node != nil {
			q = append(q, &AnnotatedNode{node: node.Left, depth: depth + 1, pos: pos << 1})
			q = append(q, &AnnotatedNode{node: node.Right, depth: depth + 1, pos: pos<<1 + 1})
			if curDepth != depth {
				curDepth = depth
				left = pos
			}
			out = Max(out, pos+1-left)
		}
	}
	return out
}

type AnnotatedNode struct {
	node  *TreeNode
	depth int
	pos   int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
