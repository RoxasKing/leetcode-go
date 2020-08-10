package leetcode

/*
  给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
  重复出现的子串要计算它们出现的次数。

  注意：
    s.length 在1到50,000之间。
    s 只包含“0”或“1”字符。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-binary-substrings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countBinarySubstrings(s string) int {
	var count int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			continue
		}
		for j, k := i, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] != s[i] || s[k] != s[i+1] {
				break
			}
			count++
		}
	}
	return count
}
