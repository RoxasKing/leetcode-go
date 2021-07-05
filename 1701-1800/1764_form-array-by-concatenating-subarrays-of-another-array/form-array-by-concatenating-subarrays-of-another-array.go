package main

func canChoose(groups [][]int, nums []int) bool {
	m, n := len(groups), len(nums)
	i, j := 0, 0
	for i < m && j < n {
		giLen := len(groups[i])
		if giLen+j <= n && isEqual(groups[i], nums[j:j+giLen]) {
			j += len(groups[i])
			i++
		} else if giLen+j < n {
			j++
		} else {
			break
		}
	}
	return i == m
}

func isEqual(nums1, nums2 []int) bool {
	for i := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
