package main

// Difficulty:
// Medium

// Tags:
// Union-Find
// Hash

func longestConsecutive(nums []int) int {
	has := map[int]struct{}{}
	for _, x := range nums {
		has[x] = struct{}{}
	}
	uf := map[int][]int{}
	var find func(x int) int
	find = func(x int) int {
		if _, ok := uf[x]; !ok {
			uf[x] = make([]int, 2)
			uf[x][0] = x
			uf[x][1] = 1
		}
		if uf[x][0] != x {
			uf[x][0] = find(uf[x][0])
		}
		return uf[x][0]
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		if uf[x][1] < uf[y][1] {
			x, y = y, x
		}
		uf[y][0] = x
		uf[x][1] += uf[y][1]
	}
	for x := range has {
		if _, ok := has[x+1]; !ok {
			continue
		}
		union(x, x+1)
	}
	o := 0
	h := map[int]int{}
	for x := range has {
		x = find(x)
		if h[x]++; o < h[x] {
			o = h[x]
		}
	}
	return o
}
