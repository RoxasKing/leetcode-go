package leetcode

func lengthOfLongestSubstring(s string) int {
	var out int
	dict := make([]int, 128)
	for i, j := 0, 0; len(s)-i >= out && j < len(s); j++ {
		i = Max(i, dict[s[j]])
		out = Max(out, j-i+1)
		dict[s[j]] = j + 1
	}
	return out
}

func longestSubstring(s string) string {
	var out string
	dict := make([]int, 128)
	for i, j := 0, 0; len(s)-i >= len(out) && j < len(s); j++ {
		i = Max(i, dict[s[j]])
		if len(out) < j-i+1 {
			out = s[i : j+1]
		}
		dict[s[j]] = j + 1
	}
	return out
}
