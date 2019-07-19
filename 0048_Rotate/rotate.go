package main

import "fmt"

func rotate(matrix [][]int) {
	out := [][]int{}

	row := len(matrix)
	col := len(matrix[0])

	for j := 0; j < col; j++ {
		tmp := []int{}
		for i := row - 1; i >= 0; i-- {
			tmp = append(tmp, matrix[i][j])
		}
		out = append(out, tmp)
	}
	for i := 0; i < len(out); i++ {
		matrix[i] = out[i]
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}
