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
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, num := range inorder {
		inorderMap[num] = i
	}
	preorderIdx := 0
	return dfs(preorder, inorderMap, &preorderIdx, 0, len(inorder)-1)
}

func dfs(preorder []int, inorderMap map[int]int, preorderIdx *int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := preorder[*preorderIdx]
	*preorderIdx++
	m := inorderMap[rootVal]
	return &TreeNode{
		Val:   rootVal,
		Left:  dfs(preorder, inorderMap, preorderIdx, l, m-1),
		Right: dfs(preorder, inorderMap, preorderIdx, m+1, r),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
