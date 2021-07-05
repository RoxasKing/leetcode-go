package main

// Tags:
// Recursion
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// BFS + Recursion
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		depth++
		curSize := len(queue)
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[curSize:]
	}
	return depth
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
