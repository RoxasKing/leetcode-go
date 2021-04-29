package main

/*
  Given two integers n and k, return the kth lexicographically smallest integer in the range [1, n].

  Example 1:
    Input: n = 13, k = 2
    Output: 10
    Explanation: The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.

  Example 2:
    Input: n = 1, k = 1
    Output: 1

  Constraints:
    1 <= k <= n <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

func findKthNumber(n int, k int) int {
	out := 1
	k--
	for k > 0 {
		step, head, tail := 0, out, out+1
		for head <= n {
			step += Min(n+1, tail) - head
			head *= 10
			tail *= 10
		}
		if step > k {
			out *= 10
			k--
		} else {
			out++
			k -= step
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
