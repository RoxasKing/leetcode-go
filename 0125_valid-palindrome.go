package leetcode

/*
  给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
  说明：本题中，我们将空字符串定义为有效的回文串。
*/

func isPalindromeII(s string) bool {
	bytes := []byte(s)
	var index int
	for i := range bytes {
		if 'a' <= bytes[i] && bytes[i] <= 'z' ||
			'0' <= bytes[i] && bytes[i] <= '9' {
			bytes[index] = bytes[i]
			index++
		} else if 'A' <= bytes[i] && bytes[i] <= 'Z' {
			bytes[index] = bytes[i] - 'A' + 'a'
			index++
		}
	}
	bytes = bytes[:index]
	for i := 0; i < len(bytes)>>1; i++ {
		if bytes[i] != bytes[len(bytes)-1-i] {
			return false
		}
	}
	return true
}
