package main

// Tags:
// Hash
func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	mark := make(map[int]bool)
	r := make(map[int]int)
	c := make(map[int]int)
	pd := make(map[int]int)
	ad := make(map[int]int)
	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		idx := x*(1e9+1) + y
		if mark[idx] {
			continue
		}
		mark[idx] = true
		r[x]++
		c[y]++
		pd[x-y+N-1]++
		ad[x+y]++
	}
	out := make([]int, len(queries))
	for i, q := range queries {
		x, y := q[0], q[1]
		out[i] = isIlluminated(N, r, c, pd, ad, x, y)
		turnOffLamps(N, mark, r, c, pd, ad, x, y)
	}
	return out
}

func turnOffLamps(N int, mark map[int]bool, r, c, pd, ad map[int]int, x, y int) {
	turnOffLamp(N, mark, r, c, pd, ad, x-1, y-1)
	turnOffLamp(N, mark, r, c, pd, ad, x-1, y)
	turnOffLamp(N, mark, r, c, pd, ad, x-1, y+1)
	turnOffLamp(N, mark, r, c, pd, ad, x, y-1)
	turnOffLamp(N, mark, r, c, pd, ad, x, y)
	turnOffLamp(N, mark, r, c, pd, ad, x, y+1)
	turnOffLamp(N, mark, r, c, pd, ad, x+1, y-1)
	turnOffLamp(N, mark, r, c, pd, ad, x+1, y)
	turnOffLamp(N, mark, r, c, pd, ad, x+1, y+1)
}

func turnOffLamp(N int, mark map[int]bool, r, c, pd, ad map[int]int, x, y int) {
	idx := x*(1e9+1) + y
	if 0 <= x && x < N && 0 <= y && y < N && mark[idx] {
		mark[idx] = false
		r[x]--
		c[y]--
		pd[x-y+N-1]--
		ad[x+y]--
	}
}

func isIlluminated(N int, r, c, pd, ad map[int]int, x, y int) int {
	if r[x] > 0 || c[y] > 0 || pd[x-y+N-1] > 0 || ad[x+y] > 0 {
		return 1
	}
	return 0
}
