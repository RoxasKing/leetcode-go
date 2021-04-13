package main

/*
  在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

  示例 1：
    输入：nums = [3,4,3,3]
    输出：4

  示例 2：
    输入：nums = [9,1,7,9,7,9,7]
    输出：1

  限制：
    1. 1 <= nums.length <= 10000
    2. 1 <= nums[i] < 2^31

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Operation
func singleNumber(nums []int) int {
	out := 0
	for i := 1; i <= 1<<31; i <<= 1 {
		cnt := 0
		for _, num := range nums {
			if num&i == i {
				cnt++
			}
		}
		if cnt%3 == 1 {
			out |= i
		}
	}
	return out
}
