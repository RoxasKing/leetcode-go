package leetcode

/*
  给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
  字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
  说明：
    字母异位词指字母相同，但排列不同的字符串。
    不考虑答案输出的顺序。
  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findAnagrams(s string, p string) []int {
	count := [26]int{}
	for _, c := range p {
		count[c-97]++
	}
	var out []int
	var l, r int
	for r < len(s) {
		count[s[r]-97]--
		for l <= r && count[s[r]-97] < 0 {
			count[s[l]-97]++
			l++
		}
		if r+1-l == len(p) {
			out = append(out, l)
			count[s[l]-97]++
			l++
		}
		r++
	}
	return out
}
