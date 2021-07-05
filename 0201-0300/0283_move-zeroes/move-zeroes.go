package main

func moveZeroes(nums []int) {
	index := 0
	for _, num := range nums {
		if num != 0 {
			nums[index] = num
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
