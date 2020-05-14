package leetcode

/*
  给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
  说明：
  你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
*/

func singleNumberII(nums []int) int {
	var seenOnce, seenTwice int
	for _, num := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}
	return seenOnce
}
