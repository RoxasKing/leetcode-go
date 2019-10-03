package _001_0025

import (
	"fmt"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	median, flag := (len1+len2)/2, (len1+len2)%2 == 0
	slice := make([]int, 2)
	flag1, flag2 := len(nums1) != 0, len(nums2) != 0
	i, j, k := 0, 0, 0
	for ; k <= median; k++ {
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

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
