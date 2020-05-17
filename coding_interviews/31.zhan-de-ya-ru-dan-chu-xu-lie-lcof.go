package codinginterviews

/*
  输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

  提示：
    0 <= pushed.length == popped.length <= 1000
    0 <= pushed[i], popped[i] < 1000
    pushed 是 popped 的排列。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func validateStackSequences(pushed []int, popped []int) bool {
	help := make([]int, 0, len(pushed))
	for len(popped) > 0 {
		for (len(help) == 0 || help[len(help)-1] != popped[0]) && len(pushed) > 0 {
			help = append(help, pushed[0])
			pushed = pushed[1:]
		}
		if help[len(help)-1] != popped[0] {
			return false
		}
		help = help[:len(help)-1]
		popped = popped[1:]
	}
	return true
}
