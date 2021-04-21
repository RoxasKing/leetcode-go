package main

/*
  Given an array A of positive integers, call a (contiguous, not necessarily distinct) subarray of A good if the number of different integers in that subarray is exactly K.
  (For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.)
  Return the number of good subarrays of A.

  Example 1:
    Input: A = [1,2,1,2,3], K = 2
    Output: 7
    Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].

  Example 2:
    Input: A = [1,2,1,3,4], K = 3
    Output: 3
    Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].

  Note:
    1. 1 <= A.length <= 20000
    2. 1 <= A[i] <= A.length
    3. 1 <= K <= A.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/subarrays-with-k-different-integers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
