package My_LeetCode_In_Go

func lengthOfLongestSubstring(s string) int {
	var out int
	dict := make([]int, 128)
	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		i = max(i, dict[s[j]])
		out = max(out, j-i+1)
		if len(s)-i < out {
			break
		}
		dict[s[j]] = j + 1
	}
	return out
}

func longestSubstring(s string) string {
	var out string
	dict := make([]int, 128)
	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		i = max(i, dict[s[j]])
		if len(out) < j-i+1 {
			out = s[i : j+1]
		}
		if len(s)-i < len(out) {
			break
		}
		dict[s[j]] = j + 1
	}
	return out
}
