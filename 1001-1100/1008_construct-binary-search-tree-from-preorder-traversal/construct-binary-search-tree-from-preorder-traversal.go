package main

// Difficulty:
// Medium

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	out := &TreeNode{Val: preorder[0]}
	i := 1
	dfs(preorder, 1<<31-1, -1<<31, &i, out)
	return out
}

func dfs(preorder []int, max, min int, i *int, node *TreeNode) {
	if *i < len(preorder) && preorder[*i] < node.Val && preorder[*i] > min {
		node.Left = &TreeNode{Val: preorder[*i]}
		*i++
		dfs(preorder, node.Val, min, i, node.Left)
	}
	if *i < len(preorder) && preorder[*i] > node.Val && preorder[*i] < max {
		node.Right = &TreeNode{Val: preorder[*i]}
		*i++
		dfs(preorder, max, node.Val, i, node.Right)
	}
}
