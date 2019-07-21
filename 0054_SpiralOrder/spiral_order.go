package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	out := []int{}
	for {
		if len(matrix) == 0 || len(matrix[0]) == 0 {
			break
		}
		// 首行顺序
		for j := 0; j < len(matrix[0]); j++ {
			out = append(out, matrix[0][j])
		}
		// 裁掉首行
		matrix = matrix[1:]
		if len(matrix) == 0 || len(matrix[0]) == 0 {
			break
		}
		// 尾列顺序
		for i := 0; i < len(matrix); i++ {
			out = append(out, matrix[i][len(matrix[i])-1])
			// 裁掉尾列
			matrix[i] = matrix[i][:len(matrix[i])-1]
		}
		if len(matrix) == 0 || len(matrix[0]) == 0 {
			break
		}
		// 尾行逆序
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			out = append(out, matrix[len(matrix)-1][j])
		}
		// 裁掉尾行
		matrix = matrix[:len(matrix)-1]
		if len(matrix) == 0 || len(matrix[0]) == 0 {
			break
		}
		// 首列逆序
		for i := len(matrix) - 1; i >= 0; i-- {
			out = append(out, matrix[i][0])
			// 裁掉首列
			matrix[i] = matrix[i][1:]
		}
	}
	return out
}

func main() {
	//in := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	in := [][]int{
		{7},
		{9},
		{6},
	}
	fmt.Println(spiralOrder(in))
}
