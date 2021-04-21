package main

/*
  给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
  例如：
  给定二叉树 [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
  返回其自底向上的层次遍历为：
    [
      [15,7],
      [9,20],
      [3]
    ]
*/

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	q := []*TreeNode{root}
	for len(q) != 0 {
		size := len(q)
		tmp := make([]int, size)
		for i := range q {
			tmp[i] = q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		out = append(out, tmp)
		q = q[size:]
	}
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
