package leetcode

/*
  给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
  例如:
  给定二叉树: [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
  返回其层次遍历结果：
    [
      [3],
      [9,20],
      [15,7]
    ]
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		cur := make([]int, 0, len(queue))
		for _, node := range queue {
			cur = append(cur, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[len(cur):]
		out = append(out, cur)
	}
	return out
}
