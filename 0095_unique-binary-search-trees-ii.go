package main

/*
  给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
*/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var generate func(int, int) []*TreeNode
	generate = func(l, r int) (res []*TreeNode) {
		if l > r {
			res = append(res, nil)
		}
		for i := l; i <= r; i++ {
			genL := generate(l, i-1)
			genR := generate(i+1, r)
			for _, l := range genL {
				for _, r := range genR {
					res = append(res, &TreeNode{Val: i, Left: l, Right: r})
				}
			}
		}
		return
	}
	return generate(1, n)
}
