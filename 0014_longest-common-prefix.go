package leetcode

/*
  编写一个函数来查找字符串数组中的最长公共前缀。
  如果不存在公共前缀，返回空字符串 ""。

  说明:
    所有输入只包含小写字母 a-z 。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ptr, ref := 0, strs[0]
Loop:
	for ptr = 0; ptr < len(ref); ptr++ {
		for _, str := range strs {
			if ptr > len(str)-1 || ref[ptr] != str[ptr] {
				break Loop
			}
		}
	}
	return ref[:ptr]
}
