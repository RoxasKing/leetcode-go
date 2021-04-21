package main

/*
  给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。
  区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

  说明:
    最直观的算法复杂度是 O(n2) ，请在此基础上优化你的算法。

  示例:
    输入: nums = [-2,5,-1], lower = -2, upper = 2,
    输出: 3
    解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-of-range-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash
func countRangeSum(nums []int, lower int, upper int) int {
	dict := make(map[int]int)
	n := len(nums)
	count, sum := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if lower <= sum && sum <= upper {
			count++
		}
		for j := lower; j <= upper; j++ {
			if v, ok := dict[sum-j]; ok {
				count += v
			}
		}
		dict[sum]++
	}
	return count
}

// Merge Sort
func countRangeSum2(nums []int, lower int, upper int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := range nums {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	prefixTmp := make([]int, n+1)
	return mergeCount(prefixSum, prefixTmp, lower, upper, 0, n)
}

func mergeCount(sum, tmp []int, lower, upper, l, r int) int {
	if l == r {
		return 0
	}

	m := l + (r-l)>>1
	count := mergeCount(sum, tmp, lower, upper, l, m) + mergeCount(sum, tmp, lower, upper, m+1, r)
	copy(tmp[l:r+1], sum[l:r+1])

	// count pair
	i, j := m+1, m+1
	for _, num := range tmp[l : m+1] {
		for i <= r && tmp[i]-num < lower {
			i++
		}
		for j <= r && tmp[j]-num <= upper {
			j++
		}
		count += j - i
	}

	// merge tmp[l:m+1] and tmp[m+1:r] to arr[l:r+1]
	p0, p1, p2 := l, l, m+1
	for p1 <= m && p2 <= r {
		if tmp[p1] < tmp[p2] {
			sum[p0] = tmp[p1]
			p1++
		} else {
			sum[p0] = tmp[p2]
			p2++
		}
		p0++
	}
	copy(sum[p0:r+1], tmp[p1:m+1])
	copy(sum[p0:r+1], tmp[p2:r+1])

	return count
}
