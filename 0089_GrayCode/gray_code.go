package main

import "fmt"

func grayCode(n int) []int {
	out := []int{}
	for i := 0; i < 1<<uint(n); i++ {
		out = append(out, i^(i>>1))
	}
	return out
}

func main() {
	n := 4
	fmt.Println(grayCode(n))
}
