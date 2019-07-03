package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	lens := len(strs)
	if lens == 0 {
		return ""
	}
	out := ""
	min := len(strs[0])
	for i := 0; i < lens; i++ {
		if strs[i] == "" {
			return ""
		}
		for j := 0; j < len(strs[i]) && j < min; j++ {
			if j == 0 && strs[0][:j+1] != strs[i][:j+1] {
				return ""
			}
			if strs[0][:j+1] == strs[i][:j+1] {
				out = strs[0][:j+1]
			}
		}
		min = len(out)
	}
	return out
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	//strs := []string{"aaa", "aa", "aa"}
	//strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
