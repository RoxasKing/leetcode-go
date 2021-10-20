package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func maxProfit(prices []int) int {
	f0, f1, f2, f3 := -1<<31, 0, -1<<31, 0
	for _, p := range prices {
		f0 = Max(f0, -p)
		f1 = Max(f1, f0+p)
		f2 = Max(f2, f1-p)
		f3 = Max(f3, f2+p)
	}
	return f3
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
