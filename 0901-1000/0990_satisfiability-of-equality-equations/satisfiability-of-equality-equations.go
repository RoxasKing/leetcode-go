package main

// Difficulty:
// Medium

// Tags:
// Union-Find

func equationsPossible(equations []string) bool {
	parent, size := [26]int{}, [26]int{}
	for i := 0; i < 26; i++ {
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
	for _, e := range equations {
		u, v := int(e[0]-'a'), int(e[3]-'a')
		if e[1] == '=' {
			union(u, v)
		}
	}
	for _, e := range equations {
		u, v := int(e[0]-'a'), int(e[3]-'a')
		if e[1] == '!' && find(u) == find(v) {
			return false
		}
	}
	return true
}
