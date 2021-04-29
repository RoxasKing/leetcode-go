package main

/*
  给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

  示例 1:
    输入: [1,2,0]
    输出: 3

  示例 2:
    输入: [3,4,-1,1]
    输出: 2

  示例 3:
    输入: [7,8,9,11,12]
    输出: 1

  提示：
    你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/first-missing-positive
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		j := nums[i] - 1
		for 0 <= j && j < n && i != j && nums[j] != nums[i] {
			nums[j], nums[i] = nums[i], nums[j]
			j = nums[i] - 1
		}
	}

	for i := range nums {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}
