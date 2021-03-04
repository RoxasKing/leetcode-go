package main

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Morris Traversal
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, last *TreeNode
	node := root
	for node != nil {
		if node.Left != nil {
			nPre := node.Left
			for nPre.Right != nil && nPre.Right != node {
				nPre = nPre.Right
			}
			if nPre.Right != node {
				nPre.Right = node
				node = node.Left
				continue
			}
		}
		if head == nil {
			head = node
		}
		if last != nil {
			last.Right = node
			node.Left = last
		}
		last = node
		node = node.Right
	}
	last.Right = head
	head.Left = last
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
