package main

import "strings"

/*
  给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

  示例：
    s = "leetcode"
    返回 0

    s = "loveleetcode"
    返回 2

  提示：你可以假定该字符串只包含小写字母。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func firstUniqChar(s string) int {
	count := [26]int{}
	for i := range s {
		count[int(s[i]-'a')]++
	}
	for i := range s {
		if count[int(s[i]-'a')] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	minIdx := len(s)
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(str); i++ {
		index := strings.IndexByte(s, str[i])
		if index != -1 && index == strings.LastIndexByte(s, str[i]) {
			minIdx = Min(minIdx, index)
		}
	}
	if minIdx == len(s) {
		return -1
	}
	return minIdx
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
