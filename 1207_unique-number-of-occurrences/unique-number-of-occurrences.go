package main

import "sort"

/*
  给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
  如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

  示例 1：
    输入：arr = [1,2,2,1,1,3]
    输出：true
    解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。

  示例 2：
    输入：arr = [1,2]
    输出：false

  示例 3：
    输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
    输出：true

  提示：
    1 <= arr.length <= 1000
    -1000 <= arr[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/unique-number-of-occurrences
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func uniqueOccurrences(arr []int) bool {
	count := map[int]int{}
	nums := []int{}
	for _, a := range arr {
		if _, ok := count[a]; !ok {
			nums = append(nums, a)
		}
		count[a]++
	}
	sort.Slice(nums, func(i, j int) bool { return count[nums[i]] < count[nums[j]] })
	for i := 1; i < len(nums); i++ {
		if count[nums[i]] == count[nums[i-1]] {
			return false
		}
	}
	return true
}

func uniqueOccurrences2(arr []int) bool {
	valCount := make([]int, 2001)
	for _, a := range arr {
		valCount[a+1000]++
	}
	countMark := make([]int, 1001)
	for _, count := range valCount {
		if count == 0 {
			continue
		}
		if countMark[count] != 0 {
			return false
		}
		countMark[count] = 1
	}
	return true
}
