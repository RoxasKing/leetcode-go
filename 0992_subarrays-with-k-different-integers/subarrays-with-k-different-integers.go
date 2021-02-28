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

// Sliding Window + Hash
func subarraysWithKDistinct(A []int, K int) int {
	out := 0
	n := len(A)
	cnt1 := make([]int, n+1)
	cnt2 := make([]int, n+1)
	l1, l2, k1, k2 := 0, 0, 0, 0

	for r := 0; r < n; r++ {
		if cnt1[A[r]] == 0 {
			k1++
		}
		if cnt2[A[r]] == 0 {
			k2++
		}
		cnt1[A[r]]++
		cnt2[A[r]]++

		for ; k1 > K; l1++ {
			cnt1[A[l1]]--
			if cnt1[A[l1]] == 0 {
				k1--
			}
		}

		for ; k2 >= K; l2++ {
			cnt2[A[l2]]--
			if cnt2[A[l2]] == 0 {
				k2--
			}
		}

		if k1 == K {
			out += l2 - l1
		}
	}

	return out
}
