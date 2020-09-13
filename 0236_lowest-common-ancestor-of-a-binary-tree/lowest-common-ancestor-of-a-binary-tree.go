package main

/*
  给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
  百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
  例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
