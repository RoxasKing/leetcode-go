package main

import "fmt"

// Tags:
// Dynamic Programming

func countSpecialSubsequences(nums []int) int {
	a, b, c := 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			a = a<<1 + 1
			a %= mod
		}
		if num == 1 {
			b = b<<1 + a
			b %= mod
		}
		if num == 2 {
			c = c<<1 + b
			c %= mod
		}
		fmt.Println(a, b, c)
	}
	return c
}

var mod = int(1e9 + 7)
