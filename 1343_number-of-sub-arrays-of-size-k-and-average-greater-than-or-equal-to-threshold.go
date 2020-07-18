package leetcode

/*
  给你一个整数数组 arr 和两个整数 k 和 threshold 。
  请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。

  提示：
    1 <= arr.length <= 10^5
    1 <= arr[i] <= 10^4
    1 <= k <= arr.length
    0 <= threshold <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numOfSubarrays(arr []int, k int, threshold int) int {
	minSum := k * threshold
	var count, sum int
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum >= minSum {
		count++
	}
	for i := k; i < len(arr); i++ {
		sum -= arr[i-k]
		sum += arr[i]
		if sum >= minSum {
			count++
		}
	}
	return count
}
