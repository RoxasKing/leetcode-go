package main

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i < m-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

// KMP
func strStr2(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	pi := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < m; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
