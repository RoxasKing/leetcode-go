package leetcode

/*
  给定一个非负整数数组，你最初位于数组的第一个位置。
  数组中的每个元素代表你在该位置可以跳跃的最大长度。
  你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*/

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var steps, nextPace int
	for i := 0; i < len(nums); i += nextPace {
		steps++
		if i+nums[i] >= len(nums)-1 {
			break
		}
		var max int // Record the maximum distance that can be reached next
		for j := 1; j <= nums[i]; j++ {
			if j+nums[i+j] > max {
				max = j + nums[i+j]
				nextPace = j
			}
		}
	}
	return steps
}

func jump2(nums []int) int {
	var steps, cur, max int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > max {
			max = nums[i] + i
		}
		if i == cur {
			steps++
			cur = max
			if cur >= len(nums)-1 {
				break
			}
		}
	}
	return steps
}
