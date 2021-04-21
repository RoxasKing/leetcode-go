package main

/*
  Given a binary array, find the maximum number of consecutive 1s in this array.

  Example 1:
    Input: [1,1,0,1,1,1]
    Output: 3
    Explanation: The first two digits or the last three digits are consecutive 1s.
        The maximum number of consecutive 1s is 3.
  Note:
    The input array will only contain 0 and 1.
    The length of input array is a positive integer and will not exceed 10,000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/max-consecutive-ones
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMaxConsecutiveOnes(nums []int) int {
	out := 0
	cnt := 0
	for _, num := range nums {
		if num == 1 {
			cnt++
			if cnt > out {
				out = cnt
			}
		} else {
			cnt = 0
		}
	}
	return out
}
