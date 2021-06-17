package main

/*
  Given an integer n, return all the numbers in the range [1, n] sorted in lexicographical order.

  You must write an algorithm that runs in O(n) time and uses O(1) extra space.

  Example 1:
    Input: n = 13
    Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]

  Example 2:
    Input: n = 2
    Output: [1,2]

  Constraints:
    1 <= n <= 5 * 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/lexicographical-numbers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func lexicalOrder(n int) []int {
	out := make([]int, n)
	for i, num := 0, 1; i < n; i++ {
		out[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			if num >= n {
				num /= 10
			}
			for num++; num%10 == 0; num /= 10 {
			}
		}
	}
	return out
}
