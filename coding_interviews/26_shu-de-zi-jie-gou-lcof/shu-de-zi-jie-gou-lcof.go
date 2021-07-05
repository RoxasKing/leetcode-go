package main

// Tags:
// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	} else if isSubTree(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubTree(A, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil || A.Val != B.Val {
		return false
	}
	return isSubTree(A.Left, B.Left) && isSubTree(A.Right, B.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
