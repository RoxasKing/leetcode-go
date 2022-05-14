package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming
// Bit Manipulation
// Bitmask

func minStickers(stickers []string, target string) int {
	n := len(target)
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	var dp func(int) int
	dp = func(mask int) int {
		if f[mask] != -1 {
			return f[mask]
		}
		f[mask] = n + 1
		for _, s := range stickers {
			left := mask
			freq := [26]int{}
			for _, c := range s {
				freq[c-'a']++
			}
			for i, c := range target {
				if (mask>>i)&1 == 1 && freq[c-'a'] > 0 {
					freq[c-'a']--
					left ^= 1 << i
				}
			}
			if left < mask && f[mask] > dp(left)+1 {
				f[mask] = f[left] + 1
			}
		}
		return f[mask]
	}
	if o := dp(1<<n - 1); o <= n {
		return o
	}
	return -1
}
