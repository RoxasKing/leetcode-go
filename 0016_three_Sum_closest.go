package My_LeetCode_In_Go

import "sort"

/*
  给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	out := nums[0] + nums[1] + nums[2]
	var i int
	for i <= len(nums)-3 {
		if nums[i] >= target/3 {
			return out
		}
		head, tail := i+1, len(nums)-1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			if sum == target {
				return sum
			} else if sum < target {
				out = closest(sum, out, target)
				for head < tail && nums[head] == nums[head+1] {
					head++
				}
				head++
			} else {
				out = closest(sum, out, target)
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				tail--
			}
		}
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return out
}

func closest(a, b, target int) int {
	t1, t2 := a-target, b-target
	if t1 < 0 {
		t1 = -t1
	}
	if t2 < 0 {
		t2 = -t2
	}
	if t1 < t2 {
		return a
	}
	return b
}
