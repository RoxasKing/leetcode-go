package leetcode

/*
  给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
*/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var generate func(int, int) []*TreeNode
	generate = func(l, r int) []*TreeNode {
		var out []*TreeNode
		if l > r {
			out = append(out, nil)
			return out
		}
		for i := l; i <= r; i++ {
			leftTrees := generate(l, i-1)
			rightTrees := generate(i+1, r)
			for _, leftTree := range leftTrees {
				for _, rightTree := range rightTrees {
					curTree := &TreeNode{
						Val:   i,
						Left:  leftTree,
						Right: rightTree,
					}
					out = append(out, curTree)
				}
			}
		}
		return out
	}
	return generate(1, n)
}
