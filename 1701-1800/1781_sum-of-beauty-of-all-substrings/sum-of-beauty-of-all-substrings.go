package main

// Tags:
// Sliding Window
func beautySum(s string) int {
	out := 0
	for k := 3; k <= len(s); k++ {
		cnt := [26]int{}
		for i := range s {
			cnt[s[i]-'a']++
			if i > k-1 {
				cnt[s[i-k]-'a']--
			}
			if i >= k-1 {
				min, max := 1<<31-1, -1<<31
				for j := 0; j < 26; j++ {
					if cnt[j] > 0 {
						min = Min(min, cnt[j])
						max = Max(max, cnt[j])
					}
				}
				out += max - min
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
