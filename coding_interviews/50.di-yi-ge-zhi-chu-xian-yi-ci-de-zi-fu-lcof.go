package codinginterviews

/*
  在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

  限制：
    0 <= s 的长度 <= 50000
*/

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	hash := [128]int{}
	for _, c := range s {
		hash[c]++
	}
	for i := range s {
		if hash[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
