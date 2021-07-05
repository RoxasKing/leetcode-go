package main

func fib(N int) int {
	dp0, dp1 := 0, 1
	for i := 1; i <= N; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp0
}
