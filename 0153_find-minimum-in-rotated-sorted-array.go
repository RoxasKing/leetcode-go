package leetcode

/*
  假设按照升序排序的数组在预先未知的某个点上进行了旋转。
  ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
  请找出其中最小的元素。
  你可以假设数组中不存在重复元素。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	if nums[l] < nums[r] {
		return nums[l]
	}
	for l < r {
		m := l + (r-l)>>1
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[m] < nums[l] {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}
