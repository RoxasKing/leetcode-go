package main

// Tags:
// Dynamic Programming

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	f0 := [50][50]int{}
	f0[startRow][startColumn] = 1
	out := 0
	for k := 0; k < maxMove; k++ {
		f1 := [50][50]int{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if f0[i][j] == 0 {
					continue
				}
				for x := 0; x < 4; x++ {
					ni, nj := i+r[x], j+c[x]
					if ni < 0 || m-1 < ni || nj < 0 || n-1 < nj {
						out += f0[i][j]
						out %= mod
						continue
					}
					f1[ni][nj] += f0[i][j]
					f1[ni][nj] %= mod
				}
			}
		}
		f0 = f1
	}
	return out
}

const mod int = 1e9 + 7

var r = [4]int{-1, 1, 0, 0}
var c = [4]int{0, 0, -1, 1}
