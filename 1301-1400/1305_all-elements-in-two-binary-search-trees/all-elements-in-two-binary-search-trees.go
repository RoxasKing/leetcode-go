package main

// Difficulty:
// Medium

// Tags:
// Morris Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	next := func(root *TreeNode) (int, *TreeNode) {
		if root == nil {
			return -1, nil
		}
		for root.Left != nil {
			pre := root.Left
			for ; pre.Right != nil && pre.Right != root; pre = pre.Right {
			}
			if pre.Right == root {
				break
			}
			pre.Right = root
			root = root.Left
		}
		return root.Val, root.Right
	}
	out := []int{}
	v1, t1 := next(root1)
	v2, t2 := next(root2)
	for v1 != -1 && v2 != -1 {
		if v1 < v2 {
			out = append(out, v1)
			v1, t1 = next(t1)
		} else {
			out = append(out, v2)
			v2, t2 = next(t2)
		}
	}
	for ; v1 != -1; v1, t1 = next(t1) {
		out = append(out, v1)
	}
	for ; v2 != -1; v2, t2 = next(t2) {
		out = append(out, v2)
	}
	return out
}
