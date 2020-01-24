package Algorithm

import (
	"fmt"
	"testing"
)

// 动态规划法解决斐波那契额数列问题
func Fibonacci(num int) int {
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

func Test_Fibonacci(t *testing.T) {
	fmt.Println(Fibonacci(1000))
}
