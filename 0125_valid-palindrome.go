package leetcode

/*
  给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
  说明：本题中，我们将空字符串定义为有效的回文串。
*/

func isPalindrome0125(s string) bool {
	array := []byte(s)
	var index int
	for i := range array {
		if 'a' <= array[i] && array[i] <= 'z' ||
			'0' <= array[i] && array[i] <= '9' {
			array[index] = array[i]
			index++
		} else if 'A' <= array[i] && array[i] <= 'Z' {
			array[index] = array[i] - 'A' + 'a'
			index++
		}
	}
	array = array[:index]
	for i := 0; i < len(array)/2; i++ {
		if array[i] != array[len(array)-1-i] {
			return false
		}
	}
	return true
}
