package codinginterviews

/*
  从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

  提示：
    节点总数 <= 1000
*/

func levelOrderII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var tmp []int
		for _, q := range queue {
			tmp = append(tmp, q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[len(tmp):]
		out = append(out, tmp)
	}
	return out
}
