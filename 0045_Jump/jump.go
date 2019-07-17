package main

import "fmt"

func jump(nums []int) int {
	size := len(nums)
	if size < 2 {
		return 0
	}
	// 步数统计
	count := 0
	// 加上下一步的最长步数
	max := 0
	// 下一步步数
	next := 0

	i := 0
	for i < size {
		// 如果最大步数超最大下标
		if i+nums[i] >= size-1 {
			return count + 1
		}
		j := 1
		for ; j <= nums[i]; j++ {
			// 如果 i+j 下标数字+当前步数比记录的更大，则更新记录
			if j+nums[i+j] >= max {
				max = j + nums[i+j]
				next = j
			}
		}
		i += next
		count++
		max = 0
		// 如果已经到了最后下标位置
		if i >= size-1 {
			return count
		}
	}

	return count
}

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{1, 2, 1, 1}
	nums := []int{2, 1}
	//nums := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println(jump(nums))
}
