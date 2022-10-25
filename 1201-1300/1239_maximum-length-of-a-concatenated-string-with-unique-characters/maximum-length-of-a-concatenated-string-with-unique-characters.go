package main

// Difficulty:
// Medium

// Tags:
// Backtracking
// Bit Manipulation

func maxLength(arr []string) int {
	o := 0
	a := make([]int, len(arr))
	for i, s := range arr {
		v := 0
		for _, c := range s {
			if v>>int(c-'a')&1 == 1 {
				v = -1
				break
			}
			v |= 1 << int(c-'a')
		}
		a[i] = v
	}
	mask, cur := 1<<26-1, 0
	var backtrack func(ptr int)
	backtrack = func(ptr int) {
		o = max(o, cur)
		if ptr == len(arr) {
			return
		}
		backtrack(ptr + 1)
		if a[ptr] == -1 || mask^a[ptr] != mask-a[ptr] {
			return
		}
		mask ^= a[ptr]
		cur += len(arr[ptr])
		backtrack(ptr + 1)
		cur -= len(arr[ptr])
		mask ^= a[ptr]
	}
	backtrack(0)
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
