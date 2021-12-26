package main

// Tags:
// Topological Sort + BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		v, u := p[0], p[1]
		edges[u] = append(edges[u], v)
		indeg[v]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		if numCourses == 0 {
			return false
		}
		numCourses--
		u := q[0]
		q = q[1:]
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return numCourses == 0
}
