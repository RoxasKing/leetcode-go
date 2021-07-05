package main

import (
	"sort"
)

// Tags:
// Two Pointers
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	out := [][]int{}
	for a := 0; a < n-3 && nums[a] <= target/4; a++ {
		for b := a + 1; b < n-2 && nums[a]+nums[b] <= target/2; b++ {
			c, d := b+1, n-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum < target {
					for c+1 < d && nums[c+1] == nums[c] {
						c++
					}
					c++
				} else if sum > target {
					for c < d-1 && nums[d-1] == nums[d] {
						d--
					}
					d--
				} else {
					out = append(out, []int{nums[a], nums[b], nums[c], nums[d]})
					for c+1 < d && nums[c+1] == nums[c] {
						c++
					}
					c++
					for c < d-1 && nums[d-1] == nums[d] {
						d--
					}
					d--
				}
			}
			for b+1 < n-2 && nums[b+1] == nums[b] {
				b++
			}
		}
		for a+1 < n-3 && nums[a+1] == nums[a] {
			a++
		}
	}
	return out
}
