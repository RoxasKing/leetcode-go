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
	cnt := make(map[int]int)
	for l, r, k := 0, 0, 0; r < len(A); r++ {
		if cnt[A[r]] == 0 {
			k++
		}
		cnt[A[r]]++
		for k > K {
			cnt[A[l]]--
			if cnt[A[l]] == 0 {
				k--
			}
			l++
		}
		if k == K {
			i := l
			for i <= r {
				out++
				cnt[A[i]]--
				if cnt[A[i]] == 0 {
					break
				}
				i++
			}
			for i >= l {
				cnt[A[i]]++
				i--
			}
		}
	}
	return out
}
