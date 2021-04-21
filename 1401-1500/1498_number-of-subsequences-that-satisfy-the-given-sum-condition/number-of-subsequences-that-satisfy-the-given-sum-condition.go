package main

import (
	"sort"
)

/*
  给你一个整数数组 nums 和一个整数 target 。
  请你统计并返回 nums 中能满足其最小元素与最大元素的 和 小于或等于 target 的 非空 子序列的数目。
  由于答案可能很大，请将结果对 10^9 + 7 取余后返回。

  提示：
    1 <= nums.length <= 10^5
    1 <= nums[i] <= 10^6
    1 <= target <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	dict := make([]int64, n)
	dict[0] = 1
	for i := int64(1); i < int64(n); i++ {
		dict[i] = (dict[i-1] * 2) % divisor
	}
	var count int64
	l, r := 0, n-1
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			count = (count + dict[r-l]) % divisor
			l++
		}
	}
	return int(count)
}

const divisor = 1e9 + 7
