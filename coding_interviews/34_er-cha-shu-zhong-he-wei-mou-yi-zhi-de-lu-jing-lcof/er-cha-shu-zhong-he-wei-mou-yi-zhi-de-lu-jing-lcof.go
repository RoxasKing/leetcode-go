package main

// Tags:
// DFS + Backtrack
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	dfs(root, sum, root.Val, []int{root.Val}, &out)
	return out
}

func dfs(node *TreeNode, sum, cur int, arr []int, out *[][]int) {
	if node.Left == nil && node.Right == nil {
		if cur == sum {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			*out = append(*out, tmp)
		}
		return
	}
	if node.Left != nil {
		cur += node.Left.Val
		arr = append(arr, node.Left.Val)
		dfs(node.Left, sum, cur, arr, out)
		arr = arr[:len(arr)-1]
		cur -= node.Left.Val
	}
	if node.Right != nil {
		cur += node.Right.Val
		arr = append(arr, node.Right.Val)
		dfs(node.Right, sum, cur, arr, out)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
