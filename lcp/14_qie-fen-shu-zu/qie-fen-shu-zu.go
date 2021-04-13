package main

/*
  给定一个整数数组 nums ，小李想将 nums 切割成若干个非空子数组，使得每个子数组最左边的数和最右边的数的最大公约数大于 1 。为了减少他的工作量，请求出最少可以切成多少个子数组。

  示例 1：
    输入：nums = [2,3,3,2,3,3]
    输出：2
    解释：最优切割为 [2,3,3,2] 和 [3,3] 。第一个子数组头尾数字的最大公约数为 2 ，第二个子数组头尾数字的最大公约数为 3 。

  示例 2：
    输入：nums = [2,3,5,7]
    输出：4
    解释：只有一种可行的切割：[2], [3], [5], [7]

  限制：
    1. 1 <= nums.length <= 10^5
    2. 2 <= nums[i] <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/qie-fen-shu-zu
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func splitArray(nums []int) int {
	n := len(nums)
	for i := 2; i < 1e6+1; i++ {
		dp[i] = n
		if minPrime[i] != 0 {
			continue
		}
		for j := i; j < 1e6+1; j += i {
			if minPrime[j] == 0 {
				minPrime[j] = i
			}
		}
	}

	out := 0
	for _, num := range nums {
		min := n
		for num > 1 {
			factor := minPrime[num]
			dp[factor] = Min(dp[factor], out+1)
			min = Min(dp[factor], min)

			num /= factor
		}
		out = min
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var minPrime [1e6 + 1]int
var dp [1e6 + 1]int
