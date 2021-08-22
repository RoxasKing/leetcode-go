package main

// DFS
// Hash

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	var out int
	dict := map[*TreeNode]int{}
	dfs(root, dict, 0, &out)
	return out % 1000000007
}

func dfs(root *TreeNode, dict map[*TreeNode]int, sum int, out *int) {
	if root == nil {
		return
	}
	sum += root.Val
	l := subTreeSum(root.Left, dict)
	r := subTreeSum(root.Right, dict)
	*out = Max(*out, Max((sum+l)*r, l*(sum+r)))
	dfs(root.Left, dict, sum+r, out)
	dfs(root.Right, dict, sum+l, out)
}

func subTreeSum(root *TreeNode, dict map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if val, ok := dict[root]; ok {
		return val
	}
	dict[root] = root.Val + subTreeSum(root.Left, dict) + subTreeSum(root.Right, dict)
	return dict[root]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
