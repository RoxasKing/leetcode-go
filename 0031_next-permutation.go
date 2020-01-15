package My_LeetCode_In_Go

/*
  实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
  如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
  必须原地修改，只允许使用额外常数空间。
  以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
    1,2,3 → 1,3,2
    3,2,1 → 1,2,3
    1,1,5 → 1,5,1
*/

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			tmp := i + 1
			for j := tmp; j <= len(nums)-2; j++ {
				if nums[j] > nums[i] && nums[j+1] <= nums[i] {
					tmp = j
					break
				}
				if j+1 == len(nums)-1 && nums[j+1] > nums[i] {
					tmp = j + 1
					break
				}
			}
			nums[i], nums[tmp] = nums[tmp], nums[i]
			break
		}
	}
	head, tail := i+1, len(nums)-1
	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head, tail = head+1, tail-1
	}
}
