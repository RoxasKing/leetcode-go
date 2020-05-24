package codinginterviews

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	var headNode, leftNode *TreeNode
	for node != nil {
		if node.Left != nil {
			preNode := node.Left
			for preNode.Right != nil && preNode.Right != node {
				preNode = preNode.Right
			}
			if preNode.Right != node {
				preNode.Right = node
				node = node.Left
				continue
			}
			node.Left = preNode
			leftNode = node
			node = node.Right
		} else {
			if headNode == nil {
				headNode = node
			}
			if leftNode != nil {
				node.Left = leftNode
			}
			leftNode = node
			node = node.Right
		}
	}
	leftNode.Right = headNode
	headNode.Left = leftNode
	return headNode
}
