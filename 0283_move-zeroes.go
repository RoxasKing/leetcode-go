package main

/*
  给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

  说明:
    必须在原数组上操作，不能拷贝额外的数组。
    尽量减少操作次数。
*/

func moveZeroes(nums []int) {
	var index int
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
