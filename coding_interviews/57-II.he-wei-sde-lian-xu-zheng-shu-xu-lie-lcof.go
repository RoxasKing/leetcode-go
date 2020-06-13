package codinginterviews

/*
  输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

  序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

  限制：
    1 <= target <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findContinuousSequence(target int) [][]int {
	var out [][]int
	l, r := 1, 2
	for l < r {
		sum := (r + 1 - l) * (l + r) >> 1
		switch {
		case sum < target:
			r++
		case sum > target:
			l++
		default:
			tmp := make([]int, r+1-l)
			for i := range tmp {
				tmp[i] = i + l
			}
			out = append(out, tmp)
			l++
		}
	}
	return out
}
