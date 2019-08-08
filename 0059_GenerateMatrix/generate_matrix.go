package main

import "fmt"

func generateMatrix(n int) [][]int {
	var out [][]int
	for i := 0; i < n; i++ {
		row := make([]int, n)
		out = append(out, row)
	}
	num := 0
	left, right, top, bot := 0, n-1, 0, n-1
	for {
		// 首行顺序
		for j := left; j <= right; j++ {
			num++
			out[top][j] = num
		}
		if num == n*n {
			break
		}
		top++
		// 尾列逆序
		for i := top; i <= bot; i++ {
			num++
			out[i][right] = num
		}
		if num == n*n {
			break
		}
		right--
		// 尾行逆序
		for j := right; j >= left; j-- {
			num++
			out[bot][j] = num
		}
		if num == n*n {
			break
		}
		bot--
		// 首列逆序
		for i := bot; i >= top; i-- {
			num++
			out[i][left] = num
		}
		if num == n*n {
			break
		}
		left++
	}
	return out
}

func main() {
	fmt.Println(generateMatrix(0))
}
