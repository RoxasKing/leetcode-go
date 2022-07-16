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

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	h := map[int]int{}
	for i, x := range inorder {
		h[x] = i
	}
	i := 0
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		val := preorder[i]
		mid := h[val]
		i++
		return &TreeNode{
			Val:   val,
			Left:  build(l, mid-1),
			Right: build(mid+1, r),
		}
	}
	return build(0, n-1)
}
