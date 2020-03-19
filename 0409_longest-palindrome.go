package leetcode

/*
  给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
  在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
  注意:
  假设字符串的长度不会超过 1010。
*/

func longestPalindrome0409(s string) int {
	count := make([]int, 128)
	for _, c := range s {
		count[c]++
	}
	var max int
	for _, v := range count {
		max += v / 2 * 2
	}
	if max < len(s) {
		max++
	}
	return max
}
