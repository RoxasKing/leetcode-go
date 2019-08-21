package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for {
		if len(matrix) == 0 {
			return false
		} else if len(matrix) == 1 {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[0][j] == target {
					return true
				}
			}
			return false
		}
		if matrix[0][len(matrix[0])-1] == target {
			return true
		} else if matrix[0][len(matrix[0])-1] < target {
			// 如果第一行最后一个数小于 target,则去掉这一行
			matrix = matrix[1:]
		} else {
			// 如果第一行最后一个数大于 target,则只保留这一行
			matrix = matrix[:1]
		}
		if len(matrix) == 0 {
			return false
		} else if len(matrix) == 1 {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[0][j] == target {
					return true
				}
			}
			return false
		}
		if matrix[len(matrix)-1][0] == target {
			return true
		} else if matrix[len(matrix)-1][0] < target {
			// 如果第一列的第一个数小于 target，则只保留这一行
			matrix = matrix[len(matrix)-1:]
		} else {
			// 如果第一列第一个数大于 target,则去掉这一行
			matrix = matrix[:len(matrix)-1]
		}
	}
}

func main() {
	//matrix := [][]int{
	//	{1, 3, 5, 7},
	//	{10, 11, 16, 20},
	//	{23, 30, 34, 50},
	//}
	//target := 3
	matrix := [][]int{
		{1},
		{3},
	}
	target := 0
	fmt.Println(searchMatrix(matrix, target))
}
