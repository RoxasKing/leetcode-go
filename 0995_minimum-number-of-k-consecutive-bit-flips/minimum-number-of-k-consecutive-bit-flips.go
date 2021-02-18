package main

/*
  In an array A containing only 0s and 1s, a K-bit flip consists of choosing a (contiguous) subarray of length K and simultaneously changing every 0 in the subarray to 1, and every 1 in the subarray to 0.

  Return the minimum number of K-bit flips required so that there is no 0 in the array.  If it is not possible, return -1.

  Example 1:
    Input: A = [0,1,0], K = 1
    Output: 2
    Explanation: Flip A[0], then flip A[2].

  Example 2:
    Input: A = [1,1,0], K = 2
    Output: -1
    Explanation: No matter how we flip subarrays of size 2, we can't make the array become [1,1,1].

  Example 3:
    Input: A = [0,0,0,1,0,1,1,0], K = 3
    Output: 3
    Explanation:
    Flip A[0],A[1],A[2]: A becomes [1,1,1,1,0,1,1,0]
    Flip A[4],A[5],A[6]: A becomes [1,1,1,1,1,0,0,0]
    Flip A[5],A[6],A[7]: A becomes [1,1,1,1,1,1,1,1]

  Note:
    1. 1 <= A.length <= 30000
    2. 1 <= K <= A.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func minKBitFlips(A []int, K int) int {
	out := 0
	a := make([]int, 0, len(A))
	b := make([]int, 0, len(A))

	for i := range A {
		a = append(a, A[i])
		b = append(b, A[i]^1)

		if i > K-1 {
			a = a[1:]
			b = b[1:]
		}

		if i >= K-1 && a[0] == 0 {
			a, b = b, a
			out++
		}
	}

	for i := range a {
		if a[i] == 0 {
			return -1
		}
	}

	return out
}
