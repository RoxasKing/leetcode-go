package main

// Tags:
// BFS
// Hash

func numBusesToDestination(routes [][]int, source int, target int) int {
	g := map[int][]int{}
	for i, arr := range routes {
		for _, route := range arr {
			g[route] = append(g[route], i)
		}
	}

	q := [][2]int{{source, 0}}
	visited := [500]bool{}
	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		cur, times := u[0], u[1]
		for _, bus := range g[cur] {
			if visited[bus] {
				continue
			}
			visited[bus] = true
			for _, route := range routes[bus] {
				if route == cur {
					continue
				}
				if route == target {
					return times + 1
				}
				q = append(q, [2]int{route, times + 1})
			}
		}
	}
	return -1
}
