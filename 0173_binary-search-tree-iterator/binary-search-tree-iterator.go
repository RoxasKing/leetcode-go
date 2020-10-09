package main

/*
  实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
  调用 next() 将返回二叉搜索树中的下一个最小的数。

  提示：
    next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
    你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-search-tree-iterator
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
