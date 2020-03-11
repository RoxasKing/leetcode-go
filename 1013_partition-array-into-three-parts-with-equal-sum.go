package leetcode

/*
  给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
  形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
  示例 1：
    输出：[0,2,1,-6,6,-7,9,1,2,0,1]
    输出：true
    解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
  示例 2：
    输入：[0,2,1,-6,6,7,9,-1,2,0,1]
    输出：false
  示例 3：
    输入：[3,3,6,5,-2,2,5,1,-9,4]
    输出：true
    解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
  提示：
    3 <= A.length <= 50000
    -10^4 <= A[i] <= 10^4
*/

func canThreePartsEqualSum(A []int) bool {
	dp := make([]int, len(A))
	dp[0] = A[0]
	for i := 1; i < len(A); i++ {
		dp[i] = dp[i-1] + A[i]
	}
	if dp[len(A)-1]%3 != 0 {
		return false
	}
	times := 1
	avg := dp[len(A)-1] / 3
	for i := 0; i < len(A)-1 && times < 3; i++ {
		if dp[i] == avg*times {
			times++
		}
	}
	return times == 3
}

func canThreePartsEqualSum2(A []int) bool {
	var sum int
	for _, a := range A {
		sum += a
	}
	if sum%3 != 0 {
		return false
	}
	avg := sum / 3
	suml, summ := 0, 0
	for i := range A {
		suml += A[i]
		if suml != avg {
			continue
		}
		for j := i + 1; j < len(A)-1; j++ {
			summ += A[j]
			if summ == avg {
				return true
			}
		}
	}
	return false
}
