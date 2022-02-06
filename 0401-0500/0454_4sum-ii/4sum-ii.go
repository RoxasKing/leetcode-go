package main

// Difficulty:
// Medium

// Tags:
// Hash

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum12 := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			sum12[num1+num2]++
		}
	}
	out := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			if cnt, ok := sum12[-(num3 + num4)]; ok {
				out += cnt
			}
		}
	}
	return out
}
