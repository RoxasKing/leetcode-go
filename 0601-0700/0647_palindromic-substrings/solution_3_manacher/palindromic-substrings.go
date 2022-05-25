package main

// DiffiCulty:
// Medium

// Tags:
// Manacher

func countSubstrings(s string) int {
	t := append(make([]byte, 0, len(s)*2+1), '#')
	for i := range s {
		t = append(t, s[i], '#')
	}
	n := len(t)
	armLen := make([]int, n)
	o := 0
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
		o += (x + 1) / 2
	}
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
