package main

import "fmt"

func setZeroes(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(matrix)
	setZeroes(matrix)
	fmt.Println(matrix)
}
