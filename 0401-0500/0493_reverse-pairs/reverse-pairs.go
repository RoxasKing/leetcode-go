package main

// Tags:
// Merge Sort
func reversePairs(nums []int) int {
	n := len(nums)
	if len(nums) < 2 {
		return 0
	}
	L := make([]int, len(nums[:n>>1]))
	copy(L, nums[:n>>1])
	R := make([]int, len(nums[n>>1:]))
	copy(R, nums[n>>1:])

	count := reversePairs(L) + reversePairs(R)
	for i, j := 0, 0; i < len(L); i++ {
		for j < len(R) && L[i] > 2*R[j] {
			j++
		}
		count += j
	}

	index := 0
	for len(L) > 0 && len(R) > 0 {
		if L[0] < R[0] {
			nums[index] = L[0]
			L = L[1:]
		} else {
			nums[index] = R[0]
			R = R[1:]
		}
		index++
	}
	copy(nums[index:], L)
	copy(nums[index:], R)
	return count
}
