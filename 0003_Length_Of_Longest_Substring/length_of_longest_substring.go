package main

import "fmt"

func lengthOfLongestSubString(s string) int {
	// 记录字符最后一次出现的位置
	location := [256]int{}
	// 初始化字符出现位置
	for i := range location {
		location[i] = -1
	}

	// 最大长度 和 起始位置
	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// 当前子串长度
		tmp := i + 1 - left
		// 当前字符最后一次出现位置如果在子串起始位置及之后
		if location[s[i]] >= left {
			// 起始位置往前移动 1 位
			left = location[s[i]] + 1
		} else if tmp > maxLen {
			maxLen = tmp
		}
		location[s[i]] = i
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubString(s))
}
