package main

/*
  Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

  Example 1:
    Input: 2
    Output: [0,1,1]

  Example 2:
    Input: 5
    Output: [0,1,1,2,1,2]

  Follow up:
    1. It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
    2. Space complexity should be O(n).
    3. Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/counting-bits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func countBits(num int) []int {
	out := make([]int, num+1)
	for i := 1; i <= num; i++ {
		out[i] = out[i>>1]
		if i&1 == 1 {
			out[i]++
		}
	}
	return out
}
