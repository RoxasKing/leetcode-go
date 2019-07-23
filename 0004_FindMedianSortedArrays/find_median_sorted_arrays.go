package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	lens := len1 + len2
	nums := make([]int, lens)
	for i := 0; i < len1; i++ {
		nums[i] = nums1[i]
	}
	for j := 0; j < len2; j++ {
		nums[len1+j] = nums2[j]
	}
	QuickSort(nums)
	if lens == 0 {
		return 0
	} else if lens%2 == 0 {
		return float64((nums[lens/2-1] + nums[lens/2])) / 2.0
	} else {
		return float64(nums[lens/2])
	}
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
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
