package main

// Tags:
// Knapsack Problem
// Dynamic Programming

func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			f[i] += f[i-coin]
		}
	}
	return f[amount]
}
