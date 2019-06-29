package main

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			tmp := i + 1
			for j := tmp; j <= len(nums)-2; j++ {
				if nums[j] > nums[i] && nums[j+1] <= nums[i] {
					tmp = j
					break
				}
				if j+1 == len(nums)-1 && nums[j+1] > nums[i] {
					tmp = j + 1
					break
				}
			}
			nums[i], nums[tmp] = nums[tmp], nums[i]
			break
		}
	}
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
