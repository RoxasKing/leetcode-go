package main

// Difficulty:
// Hard

// Tags:
// Math
// Hash
// Union-Find

func largestComponentSize(nums []int) int {
	parent := [1e5 + 1]int{}
	size := [1e5 + 1]int{}
	for i := 0; i <= 1e5; i++ {
		parent[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		if size[x] < size[y] {
			x, y = y, x
		}
		parent[y] = x
		size[x] += size[y]
	}
	exist := [1e5 + 1]bool{}
	for _, x := range nums {
		exist[x] = true
	}
	isPrime := [1e5 + 1]bool{}
	for i := 2; i <= 1e5; i++ {
		isPrime[i] = true
	}
	o := 1
	for i := 2; i <= 1e5; i++ {
		if !isPrime[i] {
			continue
		}
		for u, v := -1, i; v <= 1e5; v += i {
			isPrime[v] = false
			if !exist[v] {
				continue
			}
			if u != -1 {
				if union(u, v); o < size[find(v)] {
					o = size[find(v)]
				}
			}
			u = v
		}
	}
	return o
}
