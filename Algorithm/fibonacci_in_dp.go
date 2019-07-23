package Algorithm

// 动态规划法解决斐波那契额数列问题

// 1.自顶向下的备忘录法
func Fibonacci1(num int) int {
	if num <= 0 {
		return num
	}
	// 用于记忆斐波那契数列
	memo := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 初始化数列
		memo[i] = -1
	}
	return fib(num, memo)
}

func fib(num int, memo []int) int {
	// 如果值已经求出则直接返回
	if memo[num] != -1 {
		return memo[num]
	}
	if num <= 2 {
		// 数列前两位为 1
		memo[num] = 1
	} else {
		memo[num] = fib(num-1, memo) + fib(num-2, memo)
	}
	return memo[num]
}

// 2.自底向上的动态规划
func Fibonacci2(num int) int {
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
