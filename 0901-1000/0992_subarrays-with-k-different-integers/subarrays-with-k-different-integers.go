package main

// Tags:
// Sliding Window
func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	cntMore := make([]int, n+1)
	cntLess := make([]int, n+1)

	out := 0
	for l1, l2, r, k1, k2 := 0, 0, 0, 0, 0; r < n; r++ {
		if cntMore[A[r]] == 0 {
			k1++
		}
		if cntLess[A[r]] == 0 {
			k2++
		}
		cntMore[A[r]]++
		cntLess[A[r]]++

		for k1 > K {
			cntMore[A[l1]]--
			if cntMore[A[l1]] == 0 {
				k1--
			}
			l1++
		}

		for k2 >= K {
			cntLess[A[l2]]--
			if cntLess[A[l2]] == 0 {
				k2--
			}
			l2++
		}

		if k1 == K {
			out += l2 - l1
		}
	}
	return out
}
