package main

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	// 从倒数第二个数开始遍历,寻找到下标 i+1 大于 下标 i 的数
	for ; i >= 0; i-- {
		// 如果数 i 小于数 i+1
		if nums[i] < nums[i+1] {
			// 记录下标 i+1
			tmp := i + 1
			// 从下标 i+1 开始往后遍历，寻找到下标 j 大于 i 且 j+1 又小于等于 i 的数
			for j := tmp; j <= len(nums)-2; j++ {
				// 如果数 j 大于 数 i 且数 j+1 小于等于数 i
				if nums[j] > nums[i] && nums[j+1] <= nums[i] {
					// 获取下标 j
					tmp = j
					break
				}
				// 如果 j+1 正好是最后一个数，且下标 J+1 的数大于 i
				if j+1 == len(nums)-1 && nums[j+1] > nums[i] {
					// 获取下标 j+1
					tmp = j + 1
					break
				}
			}
			// 将上面循环获取到的最小大于 i 的数，和 i 做交换
			nums[i], nums[tmp] = nums[tmp], nums[i]
			break
		}
	}
	// 从下标 i+1 开始重新排序
	QuickSort(nums[i+1:])
}

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

func main() {
	nums := []int{5, 4, 7, 5, 5, 4, 3, 2}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
}
