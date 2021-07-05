package main

// Tags:
// Hash

func maxPoints(points [][]int) int {
	out := 1
	cntX := map[int]int{}
	cnty := map[int]int{}
	for i, p1 := range points {
		cntX[p1[0]]++
		cnty[p1[1]]++
		for j, p2 := range points[i+1:] {
			if p2[0] == p1[0] || p2[1] == p1[1] {
				continue
			}
			cnt := 2
			for _, p3 := range points[j+1:] {
				if p3[0] == p1[0] || p3[1] == p1[1] || p3[0] == p2[0] || p3[1] == p2[1] {
					continue
				}
				if (p2[0]-p1[0])*(p3[1]-p1[1]) == (p3[0]-p1[0])*(p2[1]-p1[1]) {
					cnt++
				}
			}
			out = Max(out, cnt)
		}
	}
	for _, cnt := range cntX {
		out = Max(out, cnt)
	}
	for _, cnt := range cnty {
		out = Max(out, cnt)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
