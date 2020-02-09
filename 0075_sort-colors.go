package leetcode

/*
  给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
  此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
  注意:
  不能使用代码库中的排序函数来解决这道题。
*/

func sortColors(nums []int) {
	cur, head, tail := 0, 0, len(nums)-1
	for cur <= tail {
		switch nums[cur] {
		case 0:
			nums[head], nums[cur] = nums[cur], nums[head]
			head++
			cur++
		case 1:
			cur++
		case 2:
			nums[tail], nums[cur] = nums[cur], nums[tail]
			tail--
		}
	}
}
