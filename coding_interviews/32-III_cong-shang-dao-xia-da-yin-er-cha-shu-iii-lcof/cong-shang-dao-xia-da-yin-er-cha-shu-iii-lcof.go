package main

/*
  请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

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
    [20,9],
    [15,7]
  ]

  提示：
    节点总数 <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	rev := false
	out := [][]int{}
	for len(q) > 0 {
		size := len(q)
		tmp := make([]int, 0, size)
		for _, e := range q {
			tmp = append(tmp, e.Val)
			if e.Left != nil {
				q = append(q, e.Left)
			}
			if e.Right != nil {
				q = append(q, e.Right)
			}
		}
		if rev {
			for i := 0; i < size>>1; i++ {
				tmp[i], tmp[size-1-i] = tmp[size-1-i], tmp[i]
			}
		}
		q = q[size:]
		rev = !rev
		out = append(out, tmp)
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
