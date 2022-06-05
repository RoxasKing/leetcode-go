package main

// Difficulty:
// Medium

// Tags:
// Prefix Sum
// Dynamic Programming

type NumMatrix struct {
	a [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			if i > 0 {
				val += a[i-1][j]
			}
			if j > 0 {
				val += a[i][j-1]
			}
			if i > 0 && j > 0 {
				val -= a[i-1][j-1]
			}
			a[i][j] = val
		}
	}
	return NumMatrix{a}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	a := this.a
	o := a[row2][col2]
	if row1 > 0 {
		o -= a[row1-1][col2]
	}
	if col1 > 0 {
		o -= a[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		o += a[row1-1][col1-1]
	}
	return o
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
