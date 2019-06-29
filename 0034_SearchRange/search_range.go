package main

import "fmt"

func searchRange(nums []int, target int) []int {
	size := len(nums)
	out := []int{-1, -1}

	i := 0
	for ; i < size; i++ {
		if nums[i] == target {
			out[0] = i
			break
		}
	}
	if i < size {
		for i < size-1 && nums[i+1] == target {
			i++
		}
		out[1] = i
	}
	return out
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	fmt.Println(searchRange(nums, target))
}
