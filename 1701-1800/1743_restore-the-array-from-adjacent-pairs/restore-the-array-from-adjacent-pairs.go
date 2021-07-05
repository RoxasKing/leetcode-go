package main

func restoreArray(adjacentPairs [][]int) []int {
	edges := map[int][]int{}
	for _, p := range adjacentPairs {
		u, v := p[0], p[1]
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	var u int
	for num, arr := range edges {
		if len(arr) == 1 {
			u = num
			break
		}
	}

	out := make([]int, 0, len(adjacentPairs)+1)

	for {
		out = append(out, u)
		for _, v := range edges[u] {
			if len(out) > 1 && out[len(out)-2] == v {
				continue
			}
			u = v
		}
		if u == out[len(out)-1] {
			break
		}
	}

	return out
}
