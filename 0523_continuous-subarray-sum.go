package leetcode

/*
  给定一个包含 非负数 的数组和一个目标 整数 k，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，且总和为 k 的倍数，即总和为 n*k，其中 n 也是一个整数。

  说明：
    数组的长度不会超过 10,000 。
    你可以认为所有数字总和在 32 位有符号整数范围内。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/continuous-subarray-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	remainIndex := map[int]int{0: -1}
	var sum int
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum %= k
		}
		if index, ok := remainIndex[sum]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			remainIndex[sum] = i
		}
	}
	return false
}
