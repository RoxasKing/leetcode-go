package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}
	i := 0
	for {
		if nums[i] == val {
			if len(nums) == 1 {
				return 0
			}
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
		if len(nums) == i {
			break
		}
	}
	return len(nums)
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
