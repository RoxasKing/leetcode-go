package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ss1 := strings.Split(version1, ".")
	ss2 := strings.Split(version2, ".")
	nums1 := make([]int, len(ss1))
	for i := range nums1 {
		nums1[i], _ = strconv.Atoi(ss1[i])
	}
	nums2 := make([]int, len(ss2))
	for j := range nums2 {
		nums2[j], _ = strconv.Atoi(ss2[j])
	}
	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] > nums2[0] {
			return 1
		} else if nums1[0] < nums2[0] {
			return -1
		}
		nums1 = nums1[1:]
		nums2 = nums2[1:]
	}
	for len(nums1) > 0 {
		if nums1[0] > 0 {
			return 1
		}
		nums1 = nums1[1:]
	}
	for len(nums2) > 0 {
		if nums2[0] > 0 {
			return -1
		}
		nums2 = nums2[1:]
	}
	return 0
}
