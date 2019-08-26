package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	// 检查两个字符串是否包含相同的字符
	save := make([]int, 256)
	for i := 0; i < len(s1); i++ {
		save[s1[i]]++
		save[s2[i]]--
	}
	for i := range save {
		if save[i] != 0 {
			return false
		}
	}

	// 检查字符串是否为 scramble
	for i := 1; i < len(s1); i++ {
		// 如果两个string中间某一个点，左边的substrings为scramble string，
		// 同时右边的substrings也为scramble string，则为true
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		// 如果两个string中间某一个点，s1左边的substring和s2右边的substring为scramble string,
		// 同时s1右边substring和s2左边的substring也为scramble string，则为true
		if isScramble(s1[0:i], s2[len(s1)-i:]) && isScramble(s1[i:], s2[0:len(s1)-i]) {
			return true
		}
	}
	return false
}

func main() {
	s1 := "great"
	s2 := "rgeat"
	fmt.Println(isScramble(s1, s2))
}
