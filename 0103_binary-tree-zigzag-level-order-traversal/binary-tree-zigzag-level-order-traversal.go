package main

/*
  给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

  例如：
  给定二叉树 [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
  返回锯齿形层次遍历如下：
    [
      [3],
      [20,9],
      [15,7]
    ]
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			tmp[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if reverse {
			for i := 0; i < n>>1; i++ {
				tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
			}
		}
		out = append(out, tmp)
		queue = queue[n:]
		reverse = !reverse
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
