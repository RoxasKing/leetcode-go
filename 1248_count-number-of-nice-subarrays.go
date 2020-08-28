package main

/*
  给你一个整数数组 nums 和一个整数 k。
  如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
  请返回这个数组中「优美子数组」的数目。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numberOfSubarrays(nums []int, k int) int {
	var count int
	var oddIdx []int
	for i, num := range nums {
		if num%2 == 1 {
			oddIdx = append(oddIdx, i)
		}
	}
	if len(oddIdx) < k {
		return 0
	}
	for i := 0; i <= len(oddIdx)-k; i++ {
		var leftCount, rightCount int
		if i-1 >= 0 {
			leftCount = oddIdx[i] - (oddIdx[i-1] + 1)
		} else {
			leftCount = oddIdx[i]
		}
		if i+k <= len(oddIdx)-1 {
			rightCount = oddIdx[i+k] - 1 - oddIdx[i+k-1]
		} else {
			rightCount = len(nums) - 1 - oddIdx[i+k-1]
		}
		count += 1 + leftCount + rightCount + leftCount*rightCount
	}
	return count
}
