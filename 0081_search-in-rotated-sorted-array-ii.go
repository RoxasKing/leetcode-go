package leetcode

/*
  假设按照升序排序的数组在预先未知的某个点上进行了旋转。
  ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
  编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
*/

func searchII(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	for l < r && nums[l] == nums[r] {
		l++
	}
	rotateIndex := l
	if nums[l] > nums[r] {
		rotateIndex = findRoutateIndex(&nums, l, r)
	}
	if rotateIndex == l || target < nums[0] {
		return binarySearch(&nums, rotateIndex, r, target)
	}
	return binarySearch(&nums, 0, rotateIndex, target)
}

func findRoutateIndex(nums *[]int, l, r int) int {
	for l <= r {
		m := l + (r-l)>>1
		if (*nums)[m] > (*nums)[m+1] {
			return m + 1
		}
		if (*nums)[m] >= (*nums)[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func binarySearch(nums *[]int, l, r, target int) bool {
	for l <= r {
		m := l + (r-l)>>1
		switch {
		case (*nums)[m] < target:
			l = m + 1
		case (*nums)[m] > target:
			r = m - 1
		default:
			return true
		}
	}
	return false
}
