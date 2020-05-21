package codinginterviews

/*
  输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
*/

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	var l int
	for l < len(postorder)-1 {
		if postorder[l] > root {
			break
		}
		l++
	}
	for r := l; r < len(postorder)-1; r++ {
		if postorder[r] < root {
			return false
		}
	}
	out := true
	if l > 0 {
		out = out && verifyPostorder(postorder[:l])
	}
	if l < len(postorder)-1 {
		out = out && verifyPostorder(postorder[l:len(postorder)-1])
	}
	return out
}
