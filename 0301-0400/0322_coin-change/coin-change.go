package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	f := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		f[i] = amount + 1
		for _, x := range coins {
			if x > i {
				break
			}
			if f[i] > f[i-x]+1 {
				f[i] = f[i-x] + 1
			}
		}
	}
	if f[amount] == amount+1 {
		return -1
	}
	return f[amount]
}
