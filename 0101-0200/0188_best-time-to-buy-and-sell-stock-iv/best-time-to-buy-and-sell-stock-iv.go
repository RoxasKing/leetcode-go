package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k >= n/2 {
		buy, sell := -1<<31, 0
		for _, x := range prices {
			buy = max(buy, sell-x)
			sell = max(sell, buy+x)
		}
		return sell
	}
	f := make([][2]int, k+1)
	for i := range f {
		f[i][0] = -1 << 31
	}
	for _, x := range prices {
		for i := 1; i <= k; i++ {
			f[i][0] = max(f[i][0], f[i-1][1]-x)
			f[i][1] = max(f[i][1], f[i][0]+x)
		}
	}
	return f[k][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
