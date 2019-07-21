package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	// 记录可跳跃最大长度
	max := 0
	for i := 0; i < len(nums); i++ {
		if i <= max && max > 0 || i == 0 {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		} else {
			break
		}
	}
	if max >= len(nums)-1 {
		return true
	}
	return false
}

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	nums := []int{0}
	fmt.Println(canJump(nums))
}
