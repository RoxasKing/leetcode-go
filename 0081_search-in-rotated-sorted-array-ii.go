package leetcode

/*
  假设按照升序排序的数组在预先未知的某个点上进行了旋转。
  ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
  编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
*/

func search0081(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	head, tail := 0, len(nums)-1
	for head < tail && nums[head] == nums[tail] {
		head++
	}
	rotateIndex := head
	if nums[head] > nums[tail] {
		l, r := head, tail
		for l <= r {
			m := l + (r-l)>>1
			if nums[m] > nums[m+1] {
				rotateIndex = m + 1
				break
			}
			if nums[m] < nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	binarySearch := func(l, r int) bool {
		for l <= r {
			m := l + (r-l)>>1
			switch {
			case target < nums[m]:
				r = m - 1
			case target > nums[m]:
				l = m + 1
			default:
				return true
			}
		}
		return false
	}
	if rotateIndex == head || target < nums[head] {
		return binarySearch(rotateIndex, tail)
	}
	return binarySearch(head, rotateIndex)
}
