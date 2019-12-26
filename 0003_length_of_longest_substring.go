package My_LeetCode_In_Go

func lengthOfLongestSubstring(s string) int {
	var out int
	dict := make([]int, 128)
	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		i = max(i, dict[s[j]])
		out = max(out, j-i+1)
		dict[s[j]] = j + 1
	}
	return out
}
