package main

/*
  There is an integer array perm that is a permutation of the first n positive integers, where n is always odd.

  It was encoded into another integer array encoded of length n - 1, such that encoded[i] = perm[i] XOR perm[i + 1]. For example, if perm = [1,3,2], then encoded = [2,1].

  Given the encoded array, return the original array perm. It is guaranteed that the answer exists and is unique.

  Example 1:
    Input: encoded = [3,1]
    Output: [1,2,3]
    Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]

  Example 2:
    Input: encoded = [6,5,4,6]
    Output: [2,4,1,5,3]

  Constraints:
    1. 3 <= n < 10^5
    2. n is odd.
    3. encoded.length == n - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/decode-xored-permutation
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Bit Manipulation
func decode(encoded []int) []int {
	n := len(encoded) + 1
	x := 0
	for i := 1; i <= n; i++ {
		x ^= i
	}
	for i := 1; i < n-1; i += 2 {
		x ^= encoded[i]
	}
	out := make([]int, n)
	out[0] = x
	for i := 1; i < n; i++ {
		out[i] = out[i-1] ^ encoded[i-1]
	}
	return out
}
