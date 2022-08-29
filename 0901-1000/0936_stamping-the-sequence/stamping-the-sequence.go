package main

// Difficulty:
// Hard

// Tags:
// Brute Force

func movesToStamp(stamp string, target string) []int {
	k, n := len(stamp), len(target)
	done := make([]bool, n)
	o := []int{}
	for left := n; left > 0; {
		idx, cnt := 0, 0
		for i := 0; i <= n-k; i++ {
			ok, cur := true, 0
			for j := 0; j < k; j++ {
				if done[i+j] {
					continue
				}
				if stamp[j] != target[i+j] {
					ok = false
					break
				}
				cur++
			}
			if ok && cnt < cur {
				idx, cnt = i, cur
			}
		}
		if cnt == 0 {
			return nil
		}
		for j := 0; j < k; j++ {
			if !done[idx+j] && stamp[j] == target[idx+j] {
				done[idx+j] = true
			}
		}
		left -= cnt
		o = append(o, idx)
	}
	for l, r := 0, len(o)-1; l < r; l, r = l+1, r-1 {
		o[l], o[r] = o[r], o[l]
	}
	return o
}
