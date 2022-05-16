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
	var backtrack func(int)
	backtrack = func(mask int) {
		if f[mask] != -1 {
			return
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
			if backtrack(left); f[mask] > f[left]+1 {
				f[mask] = f[left] + 1
			}
		}
	}
	if backtrack(1<<n - 1); f[1<<n-1] <= n {
		return f[1<<n-1]
	}
	return -1
}
