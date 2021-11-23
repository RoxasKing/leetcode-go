package main

// Difficulty:
// Medium

// Tags:
// Hash
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	h := map[int]int{}
	for i, v := range inorder {
		h[v] = i
	}
	i := len(postorder) - 1
	var build func(int, int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		root := &TreeNode{Val: postorder[i]}
		m := h[postorder[i]]
		i--
		root.Right = build(m+1, r)
		root.Left = build(l, m-1)
		return root
	}
	return build(0, len(postorder)-1)
}
