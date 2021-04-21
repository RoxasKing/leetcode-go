package main

/*
  Given an array of integers arr, return true if we can partition the array into three non-empty parts with equal sums.

  Formally, we can partition the array if we can find indexes i + 1 < j with (arr[0] + arr[1] + ... + arr[i] == arr[i + 1] + arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1])

  Example 1:
    Input: arr = [0,2,1,-6,6,-7,9,1,2,0,1]
    Output: true
    Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1

  Example 2:
    Input: arr = [0,2,1,-6,6,7,9,-1,2,0,1]
    Output: false

  Example 3:
    Input: arr = [3,3,6,5,-2,2,5,1,-9,4]
    Output: true
    Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4

  Constraints:
    1. 3 <= arr.length <= 5 * 10^4
    2. -10^4 <= arr[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	avg := sum / 3
	cnt := 0
	sum = 0
	for _, num := range arr {
		sum += num
		if sum == avg {
			cnt++
			sum = 0
		}
	}
	return cnt >= 3
}
