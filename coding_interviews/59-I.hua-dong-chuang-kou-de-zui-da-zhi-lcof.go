package codinginterviews

/*
  给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

  提示：
    你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
*/

// Monotonous Queue
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return nil
	}
	out := make([]int, len(nums))
	var queue []int
	for i := range nums {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		if len(queue) != 0 && i-queue[0] > k-1 {
			queue = queue[1:]
		}
		queue = append(queue, i)
		out[i] = nums[queue[0]]
	}
	return out[k-1:]
}
