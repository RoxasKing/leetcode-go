package main

/*
  给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

  数学表达式如下:
    如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
    使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
  说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/increasing-triplet-subsequence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	scd, thd := nums[len(nums)-2], nums[len(nums)-1]
	for i := len(nums) - 3; i >= 0; i-- {
		if nums[i] < scd && scd < thd {
			return true
		}
		if thd < nums[i+1] {
			scd = nums[i]
			thd = nums[i+1]
		} else if nums[i] < thd && (scd >= thd || nums[i] > scd) {
			scd = nums[i]
		}
	}
	return false
}

// Monotone Stack
func increasingTriplet2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var stack []int
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
		if len(stack) == 3 {
			return true
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
		if len(stack) == 3 {
			return true
		}
	}
	return false
}
