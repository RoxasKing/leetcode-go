package main

import "fmt"

func ExampleSolution() {
	solution := Constructor(3, 1)
	fmt.Println(solution.Flip())
	fmt.Println(solution.Flip())
	fmt.Println(solution.Flip())
	solution.Reset()
	fmt.Println(solution.Flip())
	fmt.Println(solution.Flip())
	fmt.Println(solution.Flip())
}
