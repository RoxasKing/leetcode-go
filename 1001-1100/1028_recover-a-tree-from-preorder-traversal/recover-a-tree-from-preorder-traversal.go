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
func recoverFromPreorder(S string) *TreeNode {
	i := 0
	return dfs(S, &i, 0)
}

func dfs(S string, i *int, depth int) *TreeNode {
	d := 0
	j := *i
	for j < len(S) && S[j] == '-' {
		d++
		j++
	}
	if d != depth {
		return nil
	}
	*i = j
	val := 0
	for *i < len(S) && '0' <= S[*i] && S[*i] <= '9' {
		val = val*10 + int(S[*i]-'0')
		*i++
	}
	return &TreeNode{
		Val:   val,
		Left:  dfs(S, i, depth+1),
		Right: dfs(S, i, depth+1),
	}
}

// DFS
func recoverFromPreorder2(S string) *TreeNode {
	q := []*TreeNode{}
	for S != "" {
		dep := 0
		for S[0] == '-' {
			dep++
			S = S[1:]
		}
		val := 0
		for S != "" && S[0] != '-' {
			val *= 10
			val += int(S[0] - '0')
			S = S[1:]
		}
		node := &TreeNode{Val: val}
		if dep < len(q) {
			q = q[:dep]
			q[len(q)-1].Right = node
		} else if len(q) > 0 {
			q[len(q)-1].Left = node
		}
		q = append(q, node)
	}
	return q[0]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
