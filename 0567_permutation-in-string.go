package leetcode

/*
  给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
  换句话说，第一个字符串的排列之一是第二个字符串的子串。
*/
// Sliding Window
func checkInclusion(s1 string, s2 string) bool {
	count := [26]int{}
	for _, c := range s1 {
		count[c-'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		count[s2[r]-'a']--
		for l <= r && count[s2[r]-'a'] < 0 {
			count[s2[l]-'a']++
			l++
		}
		if r+1-l == len(s1) {
			return true
		}
	}
	return false
}
