package main

// Tags:
// Topological Sorting + BFS
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	iEdges := make([][]int, n)
	iIndeg := make([]int, n)
	gEdges := make([][]int, m+n)
	gIndeg := make([]int, m+n)
	gItems := make([][]int, m+n)
	added := make(map[int]bool)

	gi := m
	for i := 0; i < n; i++ {
		if group[i] != -1 {
			continue
		}
		group[i] = gi
		gi++
	}

	for iv := range beforeItems {
		for _, iu := range beforeItems[iv] {
			iEdges[iu] = append(iEdges[iu], iv)
			iIndeg[iv]++

			gu, gv := group[iu], group[iv]
			idx := gu*(m+n+1) + gv
			if gu == gv || added[idx] {
				continue
			}
			added[idx] = true
			if added[gv*(m+n+1)+gu] {
				return []int{}
			}
			gEdges[gu] = append(gEdges[gu], gv)
			gIndeg[gv]++
		}
	}

	iq := []int{}
	for i := 0; i < n; i++ {
		if iIndeg[i] == 0 {
			iq = append(iq, i)
		}
	}

	cnt := 0
	for len(iq) > 0 {
		if cnt == n {
			return []int{}
		}
		cnt++
		u := iq[0]
		iq = iq[1:]
		gItems[group[u]] = append(gItems[group[u]], u)
		for _, v := range iEdges[u] {
			iIndeg[v]--
			if iIndeg[v] == 0 {
				iq = append(iq, v)
			}
		}
	}

	if cnt < n {
		return []int{}
	}

	gq := []int{}
	for i := 0; i < gi; i++ {
		if gIndeg[i] == 0 {
			gq = append(gq, i)
		}
	}

	out := make([]int, 0, n)
	for len(gq) > 0 {
		u := gq[0]
		gq = gq[1:]
		out = append(out, gItems[u]...)

		for _, v := range gEdges[u] {
			gIndeg[v]--
			if gIndeg[v] == 0 {
				gq = append(gq, v)
			}
		}
	}
	return out
}
