package codinginterviews

/*
  在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

  限制：
    1 <= nums.length <= 10000
    1 <= nums[i] < 2^31
*/

func singleNumber(nums []int) int {
	var out int
	for i := 1; i <= 1<<31; i <<= 1 {
		var sum int
		for _, num := range nums {
			if num&i == i {
				sum++
			}
		}
		if sum%3 == 1 {
			out |= i
		}
	}
	return out
}
