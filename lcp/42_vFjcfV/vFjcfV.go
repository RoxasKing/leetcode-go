package main

import "runtime/debug"

// Tags:
// Math
// Hash

func circleGame(toys [][]int, circles [][]int, r int) int {
	has := map[int]*state{}
	maxr, maxc := 0, 0
	for _, toy := range toys {
		x, y, i := toy[0], toy[1], toy[2]
		if i > r {
			continue
		}
		maxr = Max(maxr, x)
		maxc = Max(maxc, y)
		k := x*mod + y
		if _, ok := has[k]; !ok {
			has[k] = &state{exist: make([]bool, r+1)}
		}
		has[k].exist[i] = true
		has[k].freq++
	}
	out := 0
	solve := func(x0, y0, x, y int) {
		k := x*mod + y
		d := (x-x0)*(x-x0) + (y-y0)*(y-y0)
		for i := 1; i <= r; i++ {
			if d > (r-i)*(r-i) {
				break
			}
			if _, ok := has[k]; !ok {
				break
			}
			if has[k].exist[i] {
				out++
				has[k].exist[i] = false
				has[k].freq--
				if has[k].freq == 0 {
					delete(has, k)
				}
			}
		}
	}
	for _, circle := range circles {
		x0, y0 := circle[0], circle[1]
		solve(x0, y0, x0, y0)
		for j := 1; j < r && y0+j <= maxc; j++ {
			solve(x0, y0, x0, y0+j)
		}
		for j := -1; j > -r && y0+j >= 0; j-- {
			solve(x0, y0, x0, y0+j)
		}
		for i := 1; i < r && x0+i <= maxr; i++ {
			solve(x0, y0, x0+i, y0)
			for j := 1; j < r && y0+j <= maxc; j++ {
				solve(x0, y0, x0+i, y0+j)
			}
			for j := -1; j > -r && y0+j >= 0; j-- {
				solve(x0, y0, x0+i, y0+j)
			}
		}
		for i := -1; i > -r && x0+i >= 0; i-- {
			solve(x0, y0, x0+i, y0)
			for j := 1; j < r && y0+j <= maxc; j++ {
				solve(x0, y0, x0+i, y0+j)
			}
			for j := -1; j > -r && y0+j >= 0; j-- {
				solve(x0, y0, x0+i, y0+j)
			}
		}
	}
	return out
}

type state struct {
	exist []bool
	freq  int
}

const mod int = 1e9 + 1

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() { debug.SetGCPercent(-1) }
