package main

import "strings"

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
