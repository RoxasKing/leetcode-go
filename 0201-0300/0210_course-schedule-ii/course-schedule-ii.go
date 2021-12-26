package main

// Difficulty:
// Medium

// Tags:
// Topological Sort

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	d := make([]int, numCourses)
	for _, p := range prerequisites {
		v, u := p[0], p[1]
		g[u] = append(g[u], v)
		d[v]++
	}
	var out, q []int
	for i := 0; i < numCourses; i++ {
		if d[i] == 0 {
			q = append(q, i)
		}
	}
	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		out = append(out, u)
		for _, v := range g[u] {
			if d[v]--; d[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(out) < numCourses {
		return []int{}
	}
	return out
}
