package main

// Difficulty:
// Medium

// Tags:
// Manacher

func longestPalindrome(s string) string {
	t := append(make([]rune, 0, len(s)*2+1), '#')
	for _, c := range s {
		t = append(t, c, '#')
	}
	n := len(t)
	armLen := make([]int, n)
	var mi, mx int
	for i, k, r := 0, -1, -1; i < n; i++ {
		x := 0
		if i < r {
			x = min(armLen[k*2-i], r-i)
		}
		for ; i-x-1 >= 0 && i+x+1 < n && t[i-x-1] == t[i+x+1]; x++ {
		}
		armLen[i] = x
		if i+x > r {
			k, r = i, i+x
		}
		if mx < x {
			mi, mx = i, x
		}
	}
	o := make([]rune, 0, mx)
	for i := mi - mx + 1; i < mi+mx; i += 2 {
		o = append(o, t[i])
	}
	return string(o)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
