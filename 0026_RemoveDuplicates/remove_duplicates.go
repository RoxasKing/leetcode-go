package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i := 0
	for {
		if nums[i] == nums[i+1] {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
		if i == len(nums)-1 {
			break
		}
	}
	return len(nums)
}

func main() {
	//nums := []int{1, 1, 2, 2, 2, 3, 4}
	nums := []int{1, 2}
	fmt.Println(removeDuplicates(nums))
}
