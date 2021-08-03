package main

import "container/heap"

// Tags:
// Priority Queue
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	h := NodeHeap{}
	dfs(root, &h, 0, 0)
	out := [][]int{}
	for h.Len() > 0 {
		node := heap.Pop(&h).(Node)
		tmp := []int{node.val}
		for h.Len() > 0 && h[0].col == node.col {
			tmp = append(tmp, heap.Pop(&h).(Node).val)
		}
		out = append(out, tmp)
	}
	return out
}

func dfs(root *TreeNode, h *NodeHeap, row, col int) {
	if root == nil {
		return
	}
	heap.Push(h, Node{col: col, row: row, val: root.Val})
	dfs(root.Left, h, row+1, col-1)
	dfs(root.Right, h, row+1, col+1)
}

type Node struct {
	col, row, val int
}

type NodeHeap []Node

func (h NodeHeap) Len() int { return len(h) }
func (h NodeHeap) Less(i, j int) bool {
	if h[i].col != h[j].col {
		return h[i].col < h[j].col
	}
	if h[i].row != h[j].row {
		return h[i].row < h[j].row
	}
	return h[i].val < h[j].val
}
func (h NodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) { *h = append(*h, x.(Node)) }
func (h *NodeHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
