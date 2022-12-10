package main

// Difficulty:
// Easy

// Tags:
// Stack
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := getAllLeaf(root1)
	arr2 := getAllLeaf(root2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func getAllLeaf(root *TreeNode) []int {
	out := []int{}
	stk := []*TreeNode{root}
	for len(stk) > 0 {
		top := len(stk) - 1
		root, stk = stk[top], stk[:top]
		if root.Left == nil && root.Right == nil {
			out = append(out, root.Val)
		}
		if root.Right != nil {
			stk = append(stk, root.Right)
		}
		if root.Left != nil {
			stk = append(stk, root.Left)
		}
	}
	return out
}
