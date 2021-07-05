package main

// Tags:
// Segment Tree
// Sliding Window

func closestToTarget(arr []int, target int) int {
	n := len(arr)
	f := make([]int, 4*n)
	build(f, 1, 0, n-1, arr)

	out := 1<<31 - 1
	for l, r := 0, 0; r < n; {
		res := query(f, 1, 0, n-1, l, r)
		if res < target {
			l++
			r = Max(l, r)
		} else {
			r++
		}
		out = Min(out, Abs(res-target))
	}
	return out
}

func build(f []int, i, s, t int, arr []int) {
	if s == t {
		f[i] = arr[s]
		return
	}
	m := (s + t) >> 1
	build(f, i<<1, s, m, arr)
	build(f, i<<1+1, m+1, t, arr)
	f[i] = f[i<<1] & f[i<<1+1]
}

func query(f []int, i, s, t, l, r int) int {
	if t < l || s > r {
		return 1<<63 - 1
	}
	if l <= s && t <= r {
		return f[i]
	}
	m := (s + t) >> 1
	return query(f, i<<1, s, m, l, r) & query(f, i<<1+1, m+1, t, l, r)
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
