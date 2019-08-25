package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	num := nums[0]
	count := 1
	index := 0
	i := 1
	for {
		if i >= len(nums) {
			break
		}
		if nums[i] != num {
			num = nums[i]
			count = 1
			if index > 0 {
				copy(nums[index:], nums[i:])
				nums = nums[:len(nums)-i+index]
				i = index + 1
				index = 0
			} else {
				i++
			}
		} else {
			count++
			if count == 3 {
				index = i
			}
			i++
		}
	}
	if index > 0 {
		nums = nums[:index]
	}
	return len(nums)
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	nums = []int{1, 1, 1}
	fmt.Println(removeDuplicates(nums))
}
