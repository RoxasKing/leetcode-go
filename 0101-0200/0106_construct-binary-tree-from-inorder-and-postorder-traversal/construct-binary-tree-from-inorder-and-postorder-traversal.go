package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}
	n := len(postorder)
	pIdx := n - 1
	return dfs(inorderMap, postorder, &pIdx, 0, n-1)
}

func dfs(inorderMap map[int]int, postorder []int, pIdx *int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := postorder[*pIdx]
	*pIdx--
	mid := inorderMap[rootVal]
	return &TreeNode{
		Val:   rootVal,
		Right: dfs(inorderMap, postorder, pIdx, mid+1, r),
		Left:  dfs(inorderMap, postorder, pIdx, l, mid-1),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
