package interview

/*
  一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
*/

func singleNumbers(nums []int) []int {
	var ret int
	for _, num := range nums {
		ret ^= num
	}
	div := 1
	for div&ret == 0 {
		div <<= 1
	}
	var a, b int
	for _, num := range nums {
		if div&num == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
