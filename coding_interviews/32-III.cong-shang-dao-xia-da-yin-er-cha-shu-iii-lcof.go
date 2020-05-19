package codinginterviews

/*
  请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

  提示：
    节点总数 <= 1000
*/

func levelOrderIII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	reverse := func(nums []int) {
		for i := 0; i < len(nums)>>1; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}
	var out [][]int
	var rev bool
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
		if rev {
			reverse(tmp)
		}
		rev = !rev
		queue = queue[len(tmp):]
		out = append(out, tmp)
	}
	return out
}
