package Algorithm

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if nums[len(nums)-1] < target {
		return len(nums)
	}
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) >> 1
		if nums[mid] < target {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	return tail
}
