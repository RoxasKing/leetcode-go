package main

// Tags:
// Dynamic Programming
func minSideJumps(obstacles []int) int {
	dp0, dp1, dp2 := 1, 0, 1

	for i := 1; i < len(obstacles); i++ {
		dp00, dp01, dp02 := dp0, dp1, dp2
		switch obstacles[i] {
		case 1:
			dp0 = 1<<31 - 1
			dp1 = Min(dp01, dp02+1)
			dp2 = Min(dp02, dp01+1)
		case 2:
			dp0 = Min(dp00, dp02+1)
			dp1 = 1<<31 - 1
			dp2 = Min(dp02, dp00+1)
		case 3:
			dp0 = Min(dp00, dp01+1)
			dp1 = Min(dp01, dp00+1)
			dp2 = 1<<31 - 1
		default:
			dp0 = Min(dp00, Min(dp01+1, dp02+1))
			dp1 = Min(dp01, Min(dp00+1, dp02+1))
			dp2 = Min(dp02, Min(dp00+1, dp01+1))
		}
	}

	return Min(dp0, Min(dp1, dp2))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
