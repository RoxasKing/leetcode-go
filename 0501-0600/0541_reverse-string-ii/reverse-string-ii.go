package main

func reverseStr(s string, k int) string {
	n := len(s)
	chs := make([]byte, 0, n)
	for i := 0; i < n; i += 2 * k {
		for j := Min(i+k-1, n-1); j >= i; j-- {
			chs = append(chs, s[j])
		}
		for j := i + k; j < Min(i+2*k, n); j++ {
			chs = append(chs, s[j])
		}
	}
	return string(chs)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
