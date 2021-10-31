package main

// Difficulty:
// Medium

// Tags:
// DFS

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	_ = dfs(root, &Node{})
	root.Prev = nil
	return root
}

func dfs(cur, pre *Node) *Node {
	if cur == nil {
		return nil
	}
	for cur != nil {
		pre.Next = cur
		cur.Prev = pre
		ptr := cur
		cur = cur.Next
		if ptr.Child != nil {
			pre = dfs(ptr.Child, ptr)
			ptr.Child = nil
		} else {
			pre = ptr
		}
	}
	return pre
}
