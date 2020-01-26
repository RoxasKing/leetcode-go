package Algorithm

func Fibonacci(num int) int {
	if num <= 0 {
		return num
	}
	memo := make([]int, num+1)
	memo[0], memo[1] = 0, 1
	for i := 2; i <= num; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[num]
}
