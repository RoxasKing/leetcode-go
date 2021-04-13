package main

import "sort"

/*
  给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

  示例：
    输入：nums = [-1,2,1,-4], target = 1
    输出：2
    解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

  提示：
    3 <= nums.length <= 10^3
    -10^3 <= nums[i] <= 10^3
    -10^4 <= target <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/3sum-closest
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	out := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2 && nums[i] <= target/3; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				for l+1 < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			} else if sum > target {
				for l < r-1 && nums[r-1] == nums[r] {
					r--
				}
				r--
			} else {
				return sum
			}
			if Abs(sum-target) < Abs(out-target) {
				out = sum
			}
		}
		for i+1 < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
