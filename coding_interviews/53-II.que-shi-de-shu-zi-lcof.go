package codinginterviews

/*
  一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

  限制：
    1 <= 数组长度 <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == m {
			l = m + 1
		} else {
			if m == 0 || nums[m-1] == m-1 {
				return m
			}
			r = m - 1
		}
	}
	return l
}
