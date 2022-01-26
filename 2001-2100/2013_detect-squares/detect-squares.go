package main

// Difficulty:
// Medium

// Tags:
// Hash

type DetectSquares struct {
	h map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{h: map[int]map[int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if this.h[x] == nil {
		this.h[x] = map[int]int{}
	}
	this.h[x][y]++
}

func (this *DetectSquares) Count(point []int) int {
	out := 0
	x0, y0 := point[0], point[1]
	h0 := this.h[x0]
	for y1 := range h0 {
		if d := y1 - y0; d != 0 {
			if h1, ok := this.h[x0-d]; ok {
				out += h1[y0] * h1[y1] * h0[y1]
			}
			if h1, ok := this.h[x0+d]; ok {
				out += h1[y0] * h1[y1] * h0[y1]
			}
		}
	}
	return out
}
