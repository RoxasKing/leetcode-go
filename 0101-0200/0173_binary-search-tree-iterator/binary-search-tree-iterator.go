package main

// Difficulty:
// Medium

// Tags:
// Morris Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	t *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root}
}

func (this *BSTIterator) Next() int {
START:
	if this.t.Left != nil {
		p := this.t.Left
		for ; p.Right != nil && p.Right != this.t; p = p.Right {
		}
		if p.Right != this.t {
			p.Right = this.t
			this.t = this.t.Left
			goto START
		}
		p.Right = nil
	}
	o := this.t.Val
	this.t = this.t.Right
	return o
}

func (this *BSTIterator) HasNext() bool {
	return this.t != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
