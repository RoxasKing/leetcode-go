package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				for j+1 < k && nums[j+1] == nums[j] {
					j++
				}
				j++
			} else if sum > 0 {
				for j < k-1 && nums[k-1] == nums[k] {
					k--
				}
				k--
			} else {
				out = append(out, []int{nums[i], nums[j], nums[k]})
				for j+1 < k && nums[j+1] == nums[j] {
					j++
				}
				j++
				for j < k-1 && nums[k-1] == nums[k] {
					k--
				}
				k--
			}
		}
		for i+1 < n-2 && nums[i+1] == nums[i] {
			i++
		}
	}
	return out
}
