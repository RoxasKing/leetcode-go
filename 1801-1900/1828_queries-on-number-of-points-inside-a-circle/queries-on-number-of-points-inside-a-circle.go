package main

func countPoints(points [][]int, queries [][]int) []int {
	n := len(queries)
	out := make([]int, n)
	for i, q := range queries {
		for _, p := range points {
			if getDist(q[0], q[1], p[0], p[1]) <= q[2]*q[2] {
				out[i]++
			}
		}
	}
	return out
}

func getDist(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}
