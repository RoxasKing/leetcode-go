package main

import "container/heap"

// Difficulty:
// Hard

// Tags:
// DFS
// Priority Queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	h := minh{}
	var dfs func(node *TreeNode, x, y int)
	dfs = func(node *TreeNode, x, y int) {
		if node == nil {
			return
		}
		heap.Push(&h, pair{x, y, node.Val})
		dfs(node.Left, x+1, y-1)
		dfs(node.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
	o := [][]int{}
	top := func() int { return len(o) - 1 }
	for i := -1001; h.Len() > 0; {
		p := heap.Pop(&h).(pair)
		if i != p.y {
			i = p.y
			o = append(o, []int{})
		}
		o[top()] = append(o[top()], p.v)
	}
	return o
}

type pair struct{ x, y, v int }
type minh []pair

func (h minh) Len() int { return len(h) }
func (h minh) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.y < b.y || a.y == b.y && a.x < b.x || a.y == b.y && a.x == b.x && a.v < b.v
}
func (h minh) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minh) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minh) Pop() interface{}   { i := h.Len() - 1; o := (*h)[i]; *h = (*h)[:i]; return o }
