package main

import "sort"

// Tags:
// Monotone Chain
// Sort

func outerTrees(T [][]int) [][]int {
	n := len(T)
	O := make([][]int, 0, n<<1)
	sort.Slice(T, func(i, j int) bool { return T[i][0] < T[j][0] || T[i][0] == T[j][0] && T[i][1] < T[j][1] })
	for i := 0; i < n; i++ {
		for ; len(O) > 1 && cross(O[len(O)-2], O[len(O)-1], T[i]); O = O[:len(O)-1] {
		}
		O = append(O, T[i])
	}
	for i, m := n-2, len(O); i >= 0; i-- {
		for ; len(O) > m && cross(O[len(O)-2], O[len(O)-1], T[i]); O = O[:len(O)-1] {
		}
		O = append(O, T[i])
	}
	sort.Slice(O, func(i, j int) bool { return O[i][0] < O[j][0] || O[i][0] == O[j][0] && O[i][1] < O[j][1] })
	x := 0
	for i := 1; i < len(O); i++ {
		if O[i][0] == O[x][0] && O[i][1] == O[x][1] {
			continue
		}
		x++
		O[x] = O[i]
	}
	return O[:x+1]
}

func cross(C, A, B []int) bool { return (B[1]-C[1])*(A[0]-C[0]) < (A[1]-C[1])*(B[0]-C[0]) }
