package main

/*
  给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
  不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/

func removeDuplicates0080(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	var index int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[index] {
			if count < 2 {
				count++
				index++
				nums[index] = nums[i]
			}
		} else {
			count = 1
			index++
			nums[index] = nums[i]
		}
	}
	nums = nums[:index+1]
	return len(nums)
}
