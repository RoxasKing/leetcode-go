package leetcode

/*
  给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

  说明 :
  数组的长度为 [1, 20,000]。
  数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
*/

func subarraySum(nums []int, k int) int {
	var count, sum int
	sumMap := map[int]int{0: 1} // if sum-k == 0 , count++
	for _, num := range nums {
		sum += num
		if sumMap[sum-k] > 0 {
			count += sumMap[sum-k]
		}
		sumMap[sum]++
	}
	return count
}
