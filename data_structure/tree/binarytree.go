package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Find(data int) *TreeNode {
	for node := t; node != nil; {
		switch {
		case data < node.Val:
			node = node.Left
		case data > node.Val:
			node = node.Right
		default:
			return node
		}
	}
	return nil
}

func (t *TreeNode) Insert(data int) {
	if t == nil {
		t = &TreeNode{Val: data}
		return
	}
	for node := t; node != nil; {
		if data > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: data}
				return
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = &TreeNode{Val: data}
				return
			}
			node = node.Left
		}
	}
}

func (t *TreeNode) Delete(data int) {
	var node, nodeParent *TreeNode
	for node = t; node != nil && node.Val != data; {
		nodeParent = node
		if data < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if node == nil {
		return
	}

	if node.Left != nil && node.Right != nil {
		minNode := node.Right
		minNodeParent := node
		for minNode.Left != nil {
			minNodeParent = minNode
			minNode = minNode.Left
		}
		node.Val = minNode.Val
		node = minNode
		nodeParent = minNodeParent
	}

	var child *TreeNode
	switch {
	case node.Left != nil:
		child = node.Left
	case node.Right != nil:
		child = node.Right
	}

	switch {
	case nodeParent.Left == node:
		nodeParent.Left = child
	case nodeParent.Right == node:
		nodeParent.Right = child
	default:
		t = child
	}
}
