package main

// Tags:
// DFS + Recursion

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
	for i, v := range inorder {
		inorderMap[v] = i
	}
	i, l, r := 0, 0, len(preorder)-1
	return dfs(preorder, inorderMap, &i, l, r)
}

func dfs(preorder []int, inorderMap map[int]int, i *int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	val := preorder[*i]
	m := inorderMap[val]
	*i++
	return &TreeNode{
		Val:   val,
		Left:  dfs(preorder, inorderMap, i, l, m-1),
		Right: dfs(preorder, inorderMap, i, m+1, r),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
