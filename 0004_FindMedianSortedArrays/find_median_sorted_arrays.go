package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	// 两个数组长度的中位数  数组长度的奇偶性
	median, flag := (len1+len2)/2, (len1+len2)%2 == 0
	// 记录中位数的值
	slice := make([]int, 2)
	// 分别记录两个数组的值是否遍历完
	flag1, flag2 := len(nums1) != 0, len(nums2) != 0
	i, j, k := 0, 0, 0
	for ; k <= median; k++ {
		// 两个数组都没遍历完的情况
		if flag1 && flag2 {
			if nums1[i] < nums2[j] {
				slice[k%2] = nums1[i]
				if i < len1-1 {
					i++
				} else {
					flag1 = false
				}
			} else {
				slice[k%2] = nums2[j]
				if j < len2-1 {
					j++
				} else {
					flag2 = false
				}
			}
			// 有一个数组遍历完的情况
		} else {
			if flag1 {
				slice[k%2] = nums1[i]
				i++
			} else {
				slice[k%2] = nums2[j]
				j++
			}
		}
	}
	if flag {
		return float64(slice[0]+slice[1]) / 2.0
	} else {
		return float64(slice[(k-1)%2])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
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
	fmt.Println(findMedianSortedArrays2(nums1, nums2))
}
