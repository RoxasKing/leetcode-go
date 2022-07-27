package main

// Difficulty:
// Medium

// Tags:
// BFS
// Hash

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	sign, parent := 0, map[*TreeNode]*TreeNode{}
	for qu := []*TreeNode{root}; len(qu) != 0 && sign != 3; qu = qu[1:] {
		t := qu[0]
		if t == p {
			sign |= 1
		} else if t == q {
			sign |= 2
		}
		if t.Left != nil {
			parent[t.Left] = t
			qu = append(qu, t.Left)
		}
		if t.Right != nil {
			parent[t.Right] = t
			qu = append(qu, t.Right)
		}
	}
	vis := map[*TreeNode]struct{}{}
	for t := p; t != nil; t = parent[t] {
		vis[t] = struct{}{}
	}
	for t := q; t != nil; t = parent[t] {
		if _, ok := vis[t]; ok {
			return t
		}
	}
	return nil
}
