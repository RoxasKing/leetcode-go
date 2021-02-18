package main

/*
  Given an integer rowIndex, return the rowIndexth row of the Pascal's triangle.
  Notice that the row index starts from 0.
  In Pascal's triangle, each number is the sum of the two numbers directly above it.

  Follow up:
    Could you optimize your algorithm to use only O(k) extra space?

  Example 1:
    Input: rowIndex = 3
    Output: [1,3,3,1]

  Example 2:
    Input: rowIndex = 0
    Output: [1]

  Example 3:
    Input: rowIndex = 1
    Output: [1,1]

  Constraints:
    0 <= rowIndex <= 33

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/pascals-triangle-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func getRow(rowIndex int) []int {
	out := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		out[i] = 1
		for j := i - 1; j >= 1; j-- {
			out[j] += out[j-1]
		}
	}
	return out
}
