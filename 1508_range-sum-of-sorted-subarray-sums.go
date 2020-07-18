package leetcode

import (
	"math"
	"sort"
)

/*
  给你一个数组 nums ，它包含 n 个正整数。你需要计算所有非空连续子数组的和，并将它们按升序排序，得到一个新的包含 n * (n + 1) / 2 个数字的数组。

  请你返回在新数组中下标为 left 到 right （下标从 1 开始）的所有数字和（包括左右端点）。由于答案可能很大，请你将它对 10^9 + 7 取模后返回。

  提示：
    1 <= nums.length <= 10^3
    nums.length == n
    1 <= nums[i] <= 100
    1 <= left <= right <= n * (n + 1) / 2

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/range-sum-of-sorted-subarray-sums
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search + Double Pointer
func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	ssums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ssums[i] = ssums[i-1] + sums[i]
	}
	mod := int(math.Pow(10, 9) + 7)
	countNoBiggerSum := func(sum int) int {
		var count int
		for i, j := 0, 1; i < len(sums); i++ {
			for j < len(sums) && sums[j]-sums[i] <= sum {
				j++
			}
			count += j - 1 - i
		}
		return count
	}
	binarySearchTheKNum := func(k int) int {
		l, r := 0, sums[len(sums)-1]
		for l < r {
			m := l + (r-l)>>1
			count := countNoBiggerSum(m)
			if count < k {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}
	theKthSum := func(k int) int {
		num := binarySearchTheKNum(k)
		var preSum, preCount int
		for i, j := 0, 1; i < len(sums); i++ {
			for j < len(sums) && sums[j]-sums[i] < num {
				j++
			}
			preSum += ssums[j-1] - ssums[i] - sums[i]*(j-1-i)
			preCount += j - 1 - i
		}
		return preSum + (k-preCount)*num
	}
	return (theKthSum(right) - theKthSum(left-1) + mod) % mod
}

func rangeSum2(nums []int, n int, left int, right int) int {
	prefix := make([]int, n+1)
	for i := range nums {
		prefix[i+1] = prefix[i] + nums[i]
	}
	sons := make([]int, 0, n*(n+1))
	for i := 1; i <= len(nums); i++ {
		for j := i; j <= len(nums); j++ {
			sons = append(sons, prefix[j]-prefix[j-i])
		}
	}
	sort.Ints(sons)
	var sum int
	for i := left - 1; i < right; i++ {
		sum += sons[i]
	}
	return sum % (1e9 + 7)
}
