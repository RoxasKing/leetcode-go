package _005_Longest_Palindrome

func longestPalindrome(s string) string {
	size, out := len(s), ""
	for i, j, k := 0, 0, 0; k < size; k++ {
		if size-k-1 < len(out)/2 {
			break
		}
		i, j = k, k
		for ; i >= 0 && j < size && s[i] == s[j]; i, j = i-1, j+1 {
			if len(s[i:j+1]) > len(out) {
				out = s[i : j+1]
			}
		}
		i, j = k, k+1
		for ; i >= 0 && j < size && s[i] == s[j]; i, j = i-1, j+1 {
			if len(s[i:j+1]) > len(out) {
				out = s[i : j+1]
			}
		}
	}
	return out
}
