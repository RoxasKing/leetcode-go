package main

import "sort"

// Difficulty:
// Hard

// Tags:

type tree struct{ x, y, v int }

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	dirs := []int{-1, 0, 1, 0, -1}
	getSteps := func(x1, y1, x2, y2 int) int {
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[x1][y1] = true
		for q := []tree{{x1, y1, 0}}; len(q) > 0; q = q[1:] {
			t := q[0]
			if t.x == x2 && t.y == y2 {
				return t.v
			}
			for i := 0; i < 4; i++ {
				x, y := t.x+dirs[i], t.y+dirs[i+1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || forest[x][y] == 0 || vis[x][y] {
					continue
				}
				vis[x][y] = true
				q = append(q, tree{x, y, t.v + 1})
			}
		}
		return -1
	}
	ts := []tree{}
	for x, hs := range forest {
		for y, h := range hs {
			if h > 1 {
				ts = append(ts, tree{x, y, h})
			}
		}
	}
	sort.Slice(ts, func(i, j int) bool { return ts[i].v < ts[j].v })
	o, x, y := 0, 0, 0
	for _, t := range ts {
		v := getSteps(x, y, t.x, t.y)
		if v < 0 {
			return -1
		}
		o += v
		x, y = t.x, t.y
	}
	return o
}
