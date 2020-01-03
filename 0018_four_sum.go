package My_LeetCode_In_Go

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var out [][]int
	for i := 0; i <= len(nums)-4; {
		if nums[i] > target/4 {
			break
		}
		for j := i + 1; j <= len(nums)-3; {
			head, tail := j+1, len(nums)-1
			for head < tail {
				sum := nums[i] + nums[j] + nums[head] + nums[tail]
				if sum == target {
					out = append(
						out,
						[]int{nums[i], nums[j], nums[head], nums[tail]},
					)
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					head++
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					tail--
				} else if sum < target {
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					head++
				} else {
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					tail--
				}
			}
			for j < len(nums)-3 && nums[j] == nums[j+1] {
				j++
			}
			j++
		}
		for i < len(nums)-4 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return out
}
