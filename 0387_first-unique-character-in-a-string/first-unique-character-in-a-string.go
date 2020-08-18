package main

/*
  给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
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
