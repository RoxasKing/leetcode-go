package main

// Tags:
// Dynamic Programming

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, minProfit+1)
		f[i][0] = 1
	}

	for i := range group {
		for j := n; j >= group[i]; j-- {
			for k := 0; k <= minProfit; k++ {
				if profit[i]+k >= minProfit {
					f[j][minProfit] += f[j-group[i]][k]
					f[j][minProfit] %= 1e9 + 7
				} else {
					f[j][profit[i]+k] += f[j-group[i]][k]
					f[j][profit[i]+k] %= 1e9 + 7
				}
			}
		}
	}
	return f[n][minProfit]
}
