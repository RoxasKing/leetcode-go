package main

// Tags:
// Morris Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	node *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		node: root,
	}
}

func (this *BSTIterator) Next() int {
START:
	if this.node.Left != nil {
		prev := this.node.Left
		for prev.Right != nil && prev.Right != this.node {
			prev = prev.Right
		}
		if prev.Right != this.node {
			prev.Right = this.node
			this.node = this.node.Left
			goto START
		}
		prev.Right = nil
	}
	out := this.node.Val
	this.node = this.node.Right
	return out
}

func (this *BSTIterator) HasNext() bool {
	return this.node != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
