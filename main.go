package main

import (
	"My_LeetCode_In_Go/Algorithm"
	"fmt"
)

func main() {
	// Algorithm.Fibonacci(10)
	fmt.Println(Algorithm.Fibonacci1(10))
	fmt.Println(Algorithm.Fibonacci2(20))
	data := []int{9, 7, 23, 1, 45, 3, 11, 99, 32}
	fmt.Println(data)
	Algorithm.QuickSort(data)
	fmt.Println(data)
}
