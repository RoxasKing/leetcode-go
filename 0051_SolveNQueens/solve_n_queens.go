package main

import "fmt"

var out [][]string

func solveNQueens(n int) [][]string {
	out = [][]string{}
	rec := make(map[int]bool)
	tmp := [][]byte{}
	queen(tmp, n, 0, rec)
	return out
}

func queen(elem [][]byte, n, k int, save map[int]bool) {
	if k >= n {
		var res []string
		for _, v := range elem {
			res = append(res, string(v))
		}
		out = append(out, res)
	}
	isValid := func(elem [][]byte, i, j int) bool {
		n := 1
		for x := i - 1; x >= 0; x-- {
			if j-n >= 0 && elem[x][j-n] == 'Q' || j+n < len(elem[0]) && elem[x][j+n] == 'Q' {
				return false
			}
			n++
		}
		return true
	}
	for i := 0; i < n; i++ {
		if save[i] || !isValid(elem, k, i) {
			continue
		}
		row := []byte{}
		for j := 0; j < n; j++ {
			row = append(row, '.')
		}
		row[i] = 'Q'
		save[i] = true
		queen(append(elem, row), n, k+1, save)
		save[i] = false
	}
}

func main() {
	anss := solveNQueens(4)
	fmt.Println(anss)
}
