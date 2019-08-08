package main

import "fmt"

func generateMatrix(n int) [][]int {
	var out [][]int
	for i := 0; i < n; i++ {
		row := make([]int, n)
		out = append(out, row)
	}
	return out
}

func main() {
	fmt.Println(generateMatrix(3))
}
