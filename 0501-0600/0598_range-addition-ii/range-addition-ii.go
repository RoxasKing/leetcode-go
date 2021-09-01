package main

func maxCount(m int, n int, ops [][]int) int {
	r, c := m, n
	for _, op := range ops {
		r = Min(r, op[0])
		c = Min(c, op[1])
	}
	return r * c
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
