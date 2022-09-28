package main

// Difficulty:
// Medium

// Tags:
// Backtracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	o := [][]int{}
	a := []int{}
	sum := 0
	var backtrack func(node *TreeNode)
	backtrack = func(node *TreeNode) {
		a = append(a, node.Val)
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				t := make([]int, len(a))
				copy(t, a)
				o = append(o, t)
			}
		}
		if node.Left != nil {
			backtrack(node.Left)
		}
		if node.Right != nil {
			backtrack(node.Right)
		}
		a = a[:len(a)-1]
		sum -= node.Val
	}
	backtrack(root)
	return o
}
