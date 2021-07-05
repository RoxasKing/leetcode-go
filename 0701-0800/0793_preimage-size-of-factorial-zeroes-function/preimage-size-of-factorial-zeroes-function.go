package main

// Tags:
// Math + Binary Search
func preimageSizeFZF(k int) int {
	l, r := 0, 10*k+1
	for l < r {
		m := (l + r) >> 1
		res := countZero(m)
		if res < k {
			l = m + 1
		} else if res > k {
			r = m - 1
		} else {
			return 5
		}
	}
	return 0
}

func countZero(n int) int {
	out := 0
	for n >= 5 {
		n /= 5
		out += n
	}
	return out
}
