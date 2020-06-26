package leetcode

/*
  给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。

  在比较时，字母是依序循环出现的。举个例子：
    如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'

  提示：
    letters长度范围在[2, 10000]区间内。
    letters 仅由小写字母组成，最少包含两个不同的字母。
    目标字母target 是一个小写字母。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l <= r {
		m := l + (r-l)>>1
		if letters[m] > target {
			if m == 0 || letters[m-1] <= target {
				return letters[m]
			}
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return letters[0]
}
