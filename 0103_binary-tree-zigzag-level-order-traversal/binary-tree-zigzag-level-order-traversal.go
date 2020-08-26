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
	var out [][]int
	queue := []*TreeNode{root}
	var reverse bool
	for len(queue) != 0 {
		var tmp []int
		curSize := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if reverse {
			for i := curSize - 1; i >= 0; i-- {
				tmp = append(tmp, queue[i].Val)
			}
		} else {
			for i := 0; i < curSize; i++ {
				tmp = append(tmp, queue[i].Val)
			}
		}
		reverse = !reverse
		out = append(out, tmp)
		queue = queue[curSize:]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
