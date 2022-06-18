package main

// Difficulty:
// Medium

// Tags:
// Linked List

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		o := &Node{Val: x}
		o.Next = o
		return o
	}
	max, min := aNode.Val, aNode.Val
	last := aNode
	for p := aNode.Next; p != aNode; p = p.Next {
		if max <= p.Val {
			max, last = p.Val, p
		}
		if min >= p.Val {
			min = p.Val
		}
	}
	if min > x || max < x {
		o := &Node{Val: x}
		o.Next = last.Next
		last.Next = o
		return aNode
	}
	for p := aNode; ; p = p.Next {
		if p.Val <= x && x <= p.Next.Val {
			o := &Node{Val: x}
			o.Next = p.Next
			p.Next = o
			break
		}
	}
	return aNode
}
