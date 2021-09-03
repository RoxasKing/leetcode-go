package main

import "sort"

// Tags:
// Monotone Chain

func outerTrees(P [][]int) [][]int {
	n, x := len(P), 0
	H := make([][]int, 0, n<<1)
	sort.Slice(P, func(i, j int) bool { return P[i][0] < P[j][0] || P[i][0] == P[j][0] && P[i][1] < P[j][1] })
	for i := 0; i < n; i++ {
		for ; len(H) >= 2 && cross(H[len(H)-2], H[len(H)-1], P[i]); H = H[:len(H)-1] {
		}
		H = append(H, P[i])
	}
	for i, t := n-2, len(H)+1; i >= 0; i-- {
		for ; len(H) >= t && cross(H[len(H)-2], H[len(H)-1], P[i]); H = H[:len(H)-1] {
		}
		H = append(H, P[i])
	}
	sort.Slice(H, func(i, j int) bool { return H[i][0] < H[j][0] || H[i][0] == H[j][0] && H[i][1] < H[j][1] })
	for i := range H {
		if H[x][0] == H[i][0] && H[x][1] == H[i][1] {
			continue
		}
		x++
		H[x] = H[i]
	}
	return H[:x+1]
}

func cross(t0, t1, t2 []int) bool { return (t1[0]-t0[0])*(t2[1]-t0[1]) < (t1[1]-t0[1])*(t2[0]-t0[0]) }

// func init() { debug.SetGCPercent(-1) }
