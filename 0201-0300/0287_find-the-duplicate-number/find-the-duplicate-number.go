package main

/*
  给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

  说明：
    不能更改原数组（假设数组是只读的）。
    只能使用额外的 O(1) 的空间。
    时间复杂度小于 O(n2) 。
    数组中只有一个重复的数字，但它可能不止重复出现一次。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-the-duplicate-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}

// Binary Search
func findDuplicate2(nums []int) int {
	l, r := 1, len(nums)-1
	var out int
	for l <= r {
		m := (l + r) >> 1
		var count int
		for _, num := range nums {
			if num <= m {
				count++
			}
		}
		if count <= m {
			l = m + 1
		} else {
			r = m - 1
			out = m
		}
	}
	return out
}

// Bit operation
func findDuplicate3(nums []int) int {
	var out int
	bitMax := 31
	for (len(nums)-1)>>bitMax == 0 {
		bitMax--
	}
	for bit := 0; bit <= bitMax; bit++ {
		var x, y int
		for i, num := range nums {
			if num&(1<<bit) != 0 {
				x++
			}
			if i > 0 && i&(1<<bit) != 0 {
				y++
			}
		}
		if x > y {
			out |= 1 << bit
		}
	}
	return out
}
