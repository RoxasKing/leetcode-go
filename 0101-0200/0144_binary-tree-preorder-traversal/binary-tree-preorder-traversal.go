package main

// Tags:
// DFS
func preorderTraversal(root *TreeNode) []int {
	var out []int
	dfs(root, &out)
	return out
}

func dfs(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	*out = append(*out, root.Val)
	dfs(root.Left, out)
	dfs(root.Right, out)
}

// Stack
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := []int{}
	stk := []*TreeNode{root}
	for len(stk) > 0 {
		top := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		out = append(out, top.Val)
		if top.Right != nil {
			stk = append(stk, top.Right)
		}
		if top.Left != nil {
			stk = append(stk, top.Left)
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris Traversal
func preorderTraversal3(root *TreeNode) []int {
	out := []int{}
	node := root
	for node != nil {
		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}
			if pre.Right != node {
				pre.Right = node
				out = append(out, node.Val)
				node = node.Left
				continue
			}
			pre.Right = nil
			node = node.Right
		} else {
			out = append(out, node.Val)
			node = node.Right
		}
	}
	return out
}
