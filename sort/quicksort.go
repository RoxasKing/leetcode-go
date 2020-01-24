package sort

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := nums[0]
	head, tail := 0, len(nums)-1
	for i := 1; i <= tail; {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	QuickSort(nums[:head])
	QuickSort(nums[head+1:])
}
