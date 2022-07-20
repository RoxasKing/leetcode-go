package main

// Difficulty:
// Hard

// Tags:
// Enumeration
// BFS

func containVirus(isInfected [][]int) int {
	type pair struct{ x, y int }
	dirs := []int{-1, 0, 1, 0, -1}
	m, n := len(isInfected), len(isInfected[0])
	o := 0
	for {
		areas := []map[pair]struct{}{}
		walls := []int{}
		for i, vs := range isInfected {
			for j, v := range vs {
				if v != 1 {
					continue
				}
				q := []pair{{i, j}}
				area := map[pair]struct{}{}
				wall, idx := 0, len(areas)+1
				vs[j] = -idx
				for ; len(q) > 0; q = q[1:] {
					p := q[0]
					for k := 0; k < 4; k++ {
						x, y := p.x+dirs[k], p.y+dirs[k+1]
						if x < 0 || m-1 < x || y < 0 || n-1 < y {
							continue
						}
						if isInfected[x][y] == 1 {
							q = append(q, pair{x, y})
							isInfected[x][y] = -idx
						} else if isInfected[x][y] == 0 {
							area[pair{x, y}] = struct{}{}
							wall++
						}
					}
				}
				areas = append(areas, area)
				walls = append(walls, wall)
			}
		}
		if len(areas) == 0 {
			break
		}
		idx := 0
		for i := 1; i < len(areas); i++ {
			if len(areas[idx]) < len(areas[i]) {
				idx = i
			}
		}
		o += walls[idx]
		for _, vs := range isInfected {
			for j, v := range vs {
				if v >= 0 {
					continue
				}
				if vs[j] = 1; v == -(idx + 1) {
					vs[j]++
				}
			}
		}
		for i, area := range areas {
			if i == idx {
				continue
			}
			for p := range area {
				isInfected[p.x][p.y] = 1
			}
		}
		if len(areas) == 1 {
			break
		}
	}
	return o
}
