package main

/*
  The set [1, 2, 3, ..., n] contains a total of n! unique permutations.

  By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
    1. "123"
    2. "132"
    3. "213"
    4. "231"
    5. "312"
    6. "321"
  Given n and k, return the kth permutation sequence.

  Example 1:
    Input: n = 3, k = 3
    Output: "213"

  Example 2:
    Input: n = 4, k = 9
    Output: "2314"

  Example 3:
    Input: n = 3, k = 1
    Output: "123"

  Constraints:
    1. 1 <= n <= 9
    2. 1 <= k <= n!

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/permutation-sequence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func getPermutation(n int, k int) string {
	nums := make([]byte, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, byte(i)+'0')
	}

	k--
	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	chs := make([]byte, n)
	for i := 0; i < n-1; i++ {
		idx := k / base
		chs[i] = nums[idx]
		copy(nums[idx:], nums[idx+1:])
		k %= base
		base /= n - 1 - i
	}
	chs[n-1] = nums[0]

	return string(chs)
}
