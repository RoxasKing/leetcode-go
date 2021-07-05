package main

// Tags:
// Hungarian Algorithm
func domino(n int, m int, broken [][]int) int {
	inValid := make([]bool, n*m)
	for _, b := range broken {
		i, j := b[0], b[1]
		idx := i*m + j
		inValid[idx] = true
	}

	edges := make([][]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			idx := i*m + j
			if (i+j)&1 == 1 || inValid[idx] {
				continue
			}
			if i-1 >= 0 && !inValid[(i-1)*m+j] {
				edges[idx] = append(edges[idx], (i-1)*m+j)
			}
			if i+1 < n && !inValid[(i+1)*m+j] {
				edges[idx] = append(edges[idx], (i+1)*m+j)
			}
			if j-1 >= 0 && !inValid[i*m+j-1] {
				edges[idx] = append(edges[idx], i*m+j-1)
			}
			if j+1 < m && !inValid[i*m+j+1] {
				edges[idx] = append(edges[idx], i*m+j+1)
			}
		}
	}

	match := make([]int, n*m)
	for i := range match {
		match[i] = -1
	}

	out := 0
	for i := range edges {
		visited := make(map[int]bool)
		if find(visited, edges, match, i) {
			out++
		}
	}
	return out
}

func find(visited map[int]bool, edges [][]int, match []int, i int) bool {
	for _, j := range edges[i] {
		if visited[j] {
			continue
		}
		visited[j] = true
		if match[j] == -1 || find(visited, edges, match, match[j]) {
			match[j] = i
			return true
		}
	}
	return false
}
