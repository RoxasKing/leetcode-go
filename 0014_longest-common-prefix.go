package leetcode

/*
  编写一个函数来查找字符串数组中的最长公共前缀。
  如果不存在公共前缀，返回空字符串 ""。
*/

func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	for i := range strs {
		if len(strs[i]) == 0 {
			return ""
		}
	}
	var index int
	for index = range strs[0] {
		for i := range strs {
			if index > len(strs[i])-1 || strs[0][index] != strs[i][index] {
				return strs[0][:index]
			}
		}
	}
	return strs[0][:index+1]
}
