package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func letterCasePermutation(s string) []string {
	h := map[int]struct{}{}
	o := []string{}
	a := []byte(s)
	for i, c := range a {
		if 'A' <= c && c <= 'Z' {
			a[i] += 'a' - 'A'
		}
	}
	var mask int
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(s) {
			if _, ok := h[mask]; !ok {
				h[mask] = struct{}{}
				o = append(o, string(a))
			}
			return
		}
		backtrack(i + 1)
		if '0' <= a[i] && a[i] <= '9' {
			return
		}
		a[i] -= 'a' - 'A'
		mask ^= 1 << i
		backtrack(i + 1)
		mask ^= 1 << i
		a[i] += 'a' - 'A'
	}
	backtrack(0)
	return o
}
