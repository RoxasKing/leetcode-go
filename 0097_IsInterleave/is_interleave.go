package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	save := make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				// 空字符串必然匹配
				save[j] = true
			} else if i == 0 {
				// s1 为空字符串时，只要上一次情况匹配，且当前位置 s2 字符和 s3 对应字符相同即为匹配
				save[j] = save[j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				// s2为空字符串时，只要上一次i循环此位置匹配，且当前位置 s1 字符和 s3 对应字符相同即为匹配
				save[j] = save[j] && s1[i-1] == s3[i+j-1]
			} else {
				// s1 和 s2 都不为空时，
				// 只要 s1 第 i 个字符与 s3 第 i+j 个字符匹配, 且上一次 i 循环此位置记录为 true
				// 或者 s2 第 j 个字符与 s3 第 i+j 个字符匹配，且 j-1 时记录为 true
				save[j] = save[j] && s1[i-1] == s3[i+j-1] || save[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return save[len(s2)]
}

func main() {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}
