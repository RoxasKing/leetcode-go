package main

// Tags:
// BFS + Hash
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var flag int
	graph := make(map[*TreeNode]*TreeNode)
	queue := []*TreeNode{root}
	for len(queue) != 0 && flag != 3 {
		node := queue[0]
		queue = queue[1:]
		if node == p {
			flag |= 1
		} else if node == q {
			flag |= 2
		}
		if node.Left != nil {
			graph[node.Left] = node
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			graph[node.Right] = node
			queue = append(queue, node.Right)
		}
	}
	have := make(map[*TreeNode]bool)
	for n := p; n != nil; n = graph[n] {
		have[n] = true
	}
	for n := q; n != nil; n = graph[n] {
		if have[n] {
			return n
		}
	}
	return nil
}

// DFS + Recursion
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		} else if node == p || node == q {
			return node
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l == nil {
			return r
		} else if r == nil {
			return l
		}
		return node
	}
	return dfs(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
