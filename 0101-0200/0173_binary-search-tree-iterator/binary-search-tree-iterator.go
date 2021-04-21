package main

/*
  Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):

  BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.
  boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
  int next() Moves the pointer to the right, then returns the number at the pointer.
  Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.

  You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.

  Example 1:
    Input
      ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
      [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
    Output
      [null, 3, 7, true, 9, true, 15, true, 20, false]
    Explanation
      BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
      bSTIterator.next();    // return 3
      bSTIterator.next();    // return 7
      bSTIterator.hasNext(); // return True
      bSTIterator.next();    // return 9
      bSTIterator.hasNext(); // return True
      bSTIterator.next();    // return 15
      bSTIterator.hasNext(); // return True
      bSTIterator.next();    // return 20
      bSTIterator.hasNext(); // return False

  Constraints:
    1. The number of nodes in the tree is in the range [1, 10^5].
    2. 0 <= Node.val <= 10^6
    3. At most 10^5 calls will be made to hasNext, and next.

  Follow up:
    Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory, where h is the height of the tree?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/binary-search-tree-iterator
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
