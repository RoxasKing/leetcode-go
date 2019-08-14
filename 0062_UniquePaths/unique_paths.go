package main

import "fmt"

func uniquePaths(m int, n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return m
	}
	out := 0
	for i := m; i >= 0; i-- {
		out += i
	}
	for i := 0; i < n-3; i++ {
		out += m
		m++
	}
	return out
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
