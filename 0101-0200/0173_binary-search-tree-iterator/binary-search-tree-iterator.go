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
	return BSTIterator{node: root}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	n := this.node
	for n != nil {
		if n.Left != nil {
			pre := n.Left
			for pre.Right != nil && pre.Right != n {
				pre = pre.Right
			}
			if pre.Right != n {
				pre.Right = n
				n = n.Left
				continue
			}
			pre.Right = nil
		}
		break
	}
	out := n.Val
	this.node = n.Right
	return out
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.node != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
