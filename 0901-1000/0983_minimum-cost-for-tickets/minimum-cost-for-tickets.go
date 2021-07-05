package main

// Tags:
// Dynamic Programming
func mincostTickets(days []int, costs []int) int {
	dict := make(map[int]bool)
	for _, day := range days {
		dict[day] = true
	}
	daysMap := [3]int{1, 7, 30}
	dp := make([]int, 366)
	for i := 1; i < 366; i++ {
		if !dict[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = 1<<31 - 1
			for j, cost := range costs {
				if i-daysMap[j] > 0 {
					cost += dp[i-daysMap[j]]
				}
				dp[i] = Min(dp[i], cost)
			}
		}
	}
	return dp[365]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
