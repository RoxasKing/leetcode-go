package main

/*
  Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array. The result should also be sorted in ascending order.

  An integer a is closer to x than an integer b if:
    1. |a - x| < |b - x|, or
    2. |a - x| == |b - x| and a < b

  Example 1:
    Input: arr = [1,2,3,4,5], k = 4, x = 3
    Output: [1,2,3,4]

  Example 2:
    Input: arr = [1,2,3,4,5], k = 4, x = -1
    Output: [1,2,3,4]

  Constraints:
    1. 1 <= k <= arr.length
    2. 1 <= arr.length <= 10^4
    3. arr is sorted in ascending order.
    4. -10^4 <= arr[i], x <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-k-closest-elements
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func findClosestElements(arr []int, k int, x int) []int {
	out := make([]int, 0, k)
	for i := range arr {
		if len(out) == k && Abs(arr[i]-x) < Abs(out[0]-x) {
			out = out[1:]
			out = append(out, arr[i])
		} else if len(out) < k {
			out = append(out, arr[i])
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
