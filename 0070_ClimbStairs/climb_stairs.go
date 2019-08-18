package main

import "fmt"

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	out := make(map[int]int)
	out[0], out[1] = 1, 1
	for i := 2; i <= n; i++ {
		// 求后退一步和后退两步后的解
		out[i] = out[i-1] + out[i-2]
	}
	return out[n]
}

func main() {
	n := 10
	fmt.Println(climbStairs(n))
}
