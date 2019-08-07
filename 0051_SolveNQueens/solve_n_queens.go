package main

import "fmt"

var out [][]string

func solveNQueens(n int) [][]string {
	out = [][]string{}
	save := make(map[int]bool)
	elem := [][]byte{}
	queen(elem, n, 0, save)
	return out
}

func queen(elem [][]byte, n, i int, save map[int]bool) {
	if i >= n {
		var res []string
		for _, v := range elem {
			res = append(res, string(v))
		}
		out = append(out, res)
	}
	isValid := func(elem [][]byte, i, j int) bool {
		n := 1
		// 判断对角线上的元素是否有 ‘Q'
		for k := i - 1; k >= 0; k-- {
			if j-n >= 0 && elem[k][j-n] == 'Q' || j+n < len(elem[0]) && elem[k][j+n] == 'Q' {
				return false
			}
			n++
		}
		return true
	}
	for j := 0; j < n; j++ {
		if save[j] || !isValid(elem, i, j) {
			continue
		}
		row := []byte{}
		for k := 0; k < n; k++ {
			row = append(row, '.')
		}
		row[j] = 'Q'
		save[j] = true
		queen(append(elem, row), n, i+1, save)
		save[j] = false
	}
}

func isValid(elem [][]byte, i, j int) bool {
	n := 1
	for x := i - 1; x >= 0; x-- {
		if j-n >= 0 && elem[x][j-n] == 'Q' || j+n < len(elem[0]) && elem[x][j+n] == 'Q' {
			return false
		}
		n++
	}
	return true
}

func main() {
	fmt.Println(solveNQueens(8))
}
