package main

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	level := make([][]int, n)
	for u, ps := range preferences {
		level[u] = make([]int, n)
		for lvl, v := range ps {
			level[u][v] = lvl
		}
	}
	links := make([]int, n)
	for _, p := range pairs {
		x, y := p[0], p[1]
		links[x] = y
		links[y] = x
	}

	var out int
	for x, y := range links {
		lvl := level[x][y]
		for _, u := range preferences[x][:lvl] {
			v := links[u]
			if level[u][x] < level[u][v] {
				out++
				break
			}
		}
	}
	return out
}
