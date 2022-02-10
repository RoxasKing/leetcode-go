package main

// Difficulty:
// Hard

// Tags:
// Hash

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	turnOn := map[int]bool{}
	r := map[int]int{}
	c := map[int]int{}
	pd := map[int]int{}
	ad := map[int]int{}
	base := int(1e9 + 1)
	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		idx := x*base + y
		if turnOn[idx] {
			continue
		}
		turnOn[idx] = true
		r[x]++
		c[y]++
		pd[x-y+n-1]++
		ad[x+y]++
	}
	out := make([]int, len(queries))
	for k, q := range queries {
		x, y := q[0], q[1]
		if r[x] > 0 || c[y] > 0 || pd[x-y+n-1] > 0 || ad[x+y] > 0 {
			out[k] = 1
		}
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				idx := i*base + j
				if i < 0 || n-1 < i || j < 0 || n-1 < j || !turnOn[idx] {
					continue
				}
				turnOn[idx] = false
				r[i]--
				c[j]--
				pd[i-j+n-1]--
				ad[i+j]--
			}
		}
	}
	return out
}
