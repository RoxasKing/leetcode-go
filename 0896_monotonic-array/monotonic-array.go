package main

/*
  An array is monotonic if it is either monotone increasing or monotone decreasing.
  An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].
  Return true if and only if the given array A is monotonic.

  Example 1:
    Input: [1,2,2,3]
    Output: true

  Example 2:
    Input: [6,5,4,4]
    Output: true

  Example 3:
    Input: [1,3,2]
    Output: false

  Example 4:
    Input: [1,2,4,5]
    Output: true

  Example 5:
    Input: [1,1,1]
    Output: true

  Note:
    1. 1 <= A.length <= 50000
    2. -100000 <= A[i] <= 100000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/monotonic-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isMonotonic(A []int) bool {
	inc := []int{}
	dec := []int{}
	for _, a := range A {
		if len(inc) <= 1 {
			if len(inc) == 1 && inc[0] <= a {
				inc = inc[:0]
			}
			inc = append(inc, a)
		}
		if len(dec) <= 1 {
			if len(dec) == 1 && dec[0] >= a {
				dec = dec[:0]
			}
			dec = append(dec, a)
		}
		if len(inc) > 1 && len(dec) > 1 {
			return false
		}
	}
	return true
}
