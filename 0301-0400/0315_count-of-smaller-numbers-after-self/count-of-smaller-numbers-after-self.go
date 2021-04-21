package main

/*
  给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

  示例：
    输入：nums = [5,2,6,1]
    输出：[2,1,1,0]
    解释：
    5 的右侧有 2 个更小的元素 (2 和 1)
    2 的右侧仅有 1 个更小的元素 (1)
    6 的右侧有 1 个更小的元素 (1)
    1 的右侧有 0 个更小的元素

  提示：
    0 <= nums.length <= 10^5
    -10^4 <= nums[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Merge Sort
func countSmaller(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	counter := make([]int, n)
	tmpIdxs := make([]int, n)
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	merge(nums, counter, tmpIdxs, indexes, 0, n-1)
	return counter
}

func merge(nums, counter, tmpIdxs, indexes []int, l, r int) {
	if l == r {
		return
	}
	m := l + (r-l)>>1
	merge(nums, counter, tmpIdxs, indexes, l, m)
	merge(nums, counter, tmpIdxs, indexes, m+1, r)
	if nums[indexes[m]] <= nums[indexes[m+1]] {
		return
	}
	copy(tmpIdxs[l:r+1], indexes[l:r+1])
	i, j, k := l, m+1, l
	for i <= m && j <= r {
		if nums[tmpIdxs[i]] > nums[tmpIdxs[j]] {
			indexes[k] = tmpIdxs[j]
			j++
		} else {
			counter[tmpIdxs[i]] += j - 1 - m
			indexes[k] = tmpIdxs[i]
			i++
		}
		k++
	}
	copy(indexes[k:k+r+1-j], tmpIdxs[j:r+1])
	copy(indexes[k:k+m+1-i], tmpIdxs[i:m+1])
	for i <= m {
		counter[tmpIdxs[i]] += r - m
		i++
		k++
	}
}
