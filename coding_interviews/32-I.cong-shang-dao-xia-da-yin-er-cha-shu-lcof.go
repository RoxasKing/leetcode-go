package codinginterviews

/*
  从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

  提示：
    节点总数 <= 1000
*/

// Preorder traversal (VLR) + BFS
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var out []int
	q := []*TreeNode{root}
	for len(q) != 0 {
		root, q = q[0], q[1:]
		out = append(out, root.Val)
		if root.Left != nil {
			q = append(q, root.Left)
		}
		if root.Right != nil {
			q = append(q, root.Right)
		}
	}
	return out
}
