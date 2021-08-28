package main

// Tags:
// Greedy

func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	cnt1, cnt2 := [7]int{}, [7]int{}
	for _, num := range nums1 {
		sum1 += num
		cnt1[num]++
	}
	for _, num := range nums2 {
		sum2 += num
		cnt2[num]++
	}

	out := 0
	for sum1 != sum2 {
		out++

		if sum1 < sum2 {
			sum1, sum2 = sum2, sum1
			cnt1, cnt2 = cnt2, cnt1
		}
		diff := sum1 - sum2

		dec, decIdx := -1, -1
		for i := 6; i > 1; i-- {
			if cnt1[i] > 0 {
				dec, decIdx = i-1, i
				if i-diff > 1 {
					dec = diff
				}
				break
			}
		}

		if dec == diff {
			break
		}

		inc, incIdx := -1, -1
		for i := 1; i < 6; i++ {
			if cnt2[i] > 0 {
				inc, incIdx = 6-i, i
				if i+diff < 6 {
					inc = diff
				}
				break
			}
		}

		if inc == diff {
			break
		}

		if decIdx == -1 && incIdx == -1 {
			return -1
		} else if decIdx != -1 && incIdx != -1 {
			if dec > inc {
				sum1 -= dec
				cnt1[decIdx]--
				cnt1[decIdx-dec]++
			} else {
				sum2 += inc
				cnt2[incIdx]--
				cnt2[incIdx+inc]++
			}
		} else if decIdx != -1 {
			sum1 -= dec
			cnt1[decIdx]--
			cnt1[decIdx-dec]++
		} else if incIdx != -1 {
			sum2 += inc
			cnt2[incIdx]--
			cnt2[incIdx+inc]++
		}

	}
	return out
}
