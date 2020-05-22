package codinginterviews

/*
  输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
*/

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	var verify func(int, int) bool
	verify = func(l, r int) bool {
		if l >= r {
			return true
		}
		rootVal := postorder[r]
		index := l
		for index < r && postorder[index] <= rootVal {
			index++
		}
		for i := index; i < r; i++ {
			if postorder[i] < rootVal {
				return false
			}
		}
		return verify(l, index-1) && verify(index, r-1)
	}
	return verify(0, len(postorder)-1)
}
