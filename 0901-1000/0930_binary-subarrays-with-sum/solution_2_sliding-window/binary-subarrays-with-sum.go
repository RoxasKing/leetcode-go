package main

// Tags:
// Sliding Window

func numSubarraysWithSum(nums []int, goal int) int {
	var out int
	for l1, l2, r, cnt1, cnt2 := 0, 0, 0, 0, 0; r < len(nums); r++ {
		for cnt1 += nums[r]; l1 <= r && cnt1 > goal; l1++ {
			cnt1 -= nums[l1]
		}
		for cnt2 += nums[r]; l2 <= r && cnt2 >= goal; l2++ {
			cnt2 -= nums[l2]
		}
		out += l2 - l1
	}
	return out
}
