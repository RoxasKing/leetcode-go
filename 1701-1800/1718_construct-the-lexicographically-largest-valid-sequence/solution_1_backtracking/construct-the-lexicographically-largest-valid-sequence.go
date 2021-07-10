package main

// Tags:
// Backtracking

func constructDistancedSequence(n int) []int {
	out := make([]int, (n-1)*2+1)
	cur := make([]int, (n-1)*2+1)
	backtrack(out, cur, (n-1)*2+1, n)
	return out
}

func backtrack(out, cur []int, n int, x int) {
	for i := 0; i < n; i++ {
		if cur[i] > out[i] || cur[i] == 0 && x > out[i] {
			break
		} else if cur[i] == 0 && x < out[i] {
			return
		}
	}

	if x == 1 {
		for i := range cur {
			if cur[i] > 0 {
				continue
			}
			cur[i] = 1
			if bigger(cur, out) {
				copy(out, cur)
			}
			cur[i] = 0
			break
		}
		return
	}

	for i := 0; i+x < n; i++ {
		if cur[i] != 0 || cur[i+x] != 0 {
			continue
		}
		cur[i] = x
		cur[i+x] = x
		backtrack(out, cur, n, x-1)
		cur[i] = 0
		cur[i+x] = 0
	}
}

func bigger(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return false
}
