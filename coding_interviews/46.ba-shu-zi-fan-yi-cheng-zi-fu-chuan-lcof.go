package codinginterviews

import "strconv"

/*
  给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

  提示：
    0 <= num < 2^31

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Recursive
func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	var count int
	var recur func(int)
	recur = func(index int) {
		if index == len(numStr) {
			count++
			return
		}
		recur(index + 1)
		if index+1 < len(numStr) &&
			(numStr[index] == '1' || numStr[index] == '2' && numStr[index+1] <= '5') {
			recur(index + 2)
		}
	}
	recur(0)
	return count
}

// Dynamic Programming
func translateNum2(num int) int {
	numStr := strconv.Itoa(num)
	if len(numStr) == 1 {
		return 1
	}
	dp1, dp2 := 1, 1
	for i := 1; i < len(numStr); i++ {
		if numStr[i-1] == '1' || numStr[i-1] == '2' && numStr[i] <= '5' {
			dp1, dp2 = dp2, dp1+dp2
		} else {
			dp1 = dp2
		}
	}
	return dp2
}
