package main

// Tags:
// Sliding Window
func numberOfSubarrays(nums []int, k int) int {
	oddIdx := []int{-1}
	for i := range nums {
		if nums[i]&1 == 1 {
			oddIdx = append(oddIdx, i)
		}
	}
	oddIdx = append(oddIdx, len(nums))

	if len(oddIdx)-2 < k {
		return 0
	}

	out := 0
	for i := 1; i <= len(oddIdx)-1-k; i++ {
		lCnt := oddIdx[i] - (oddIdx[i-1] + 1)
		rCnt := oddIdx[i+k] - 1 - oddIdx[i+k-1]
		out += 1 + lCnt + rCnt + lCnt*rCnt
	}
	return out
}

// Two Pointers + Sliding Window
func numberOfSubarrays2(nums []int, k int) int {
	odd := 0
	out := 0
	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r]&1 == 1 {
			odd++
		}
		for l <= r && odd > k {
			if nums[l]&1 == 1 {
				odd--
			}
			l++
		}
		if odd == k {
			for i := l; i <= r; i++ {
				out++
				if nums[i]&1 == 1 {
					break
				}
			}
		}
	}
	return out
}
