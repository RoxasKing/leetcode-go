package main

// 动态规划法解决斐波那契额数列问题

import "fmt"

// 1.自顶向下的备忘录法
func fibonacci1(num int) int {
	if num <= 0 {
		return num
	}
	memo := make([]int, num+1)
	for i := 0; i <= num; i++ {
		memo[i] = -1
	}
	return fib(num, memo)
}

func fib(num int, memo []int) int {
	if memo[num] != -1 {
		return memo[num]
	}
	if num <= 2 {
		memo[num] = 1
	} else {
		memo[num] = fib(num-1, memo) + fib(num-2, memo)
	}
	return memo[num]
}

func fibonacci2(num int) int {
	if num <= 0 {
		return num
	}
	memo := make([]int, num+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= num; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[num]
}

func main() {
	num := 10
	fmt.Println(fibonacci2(num))
}
